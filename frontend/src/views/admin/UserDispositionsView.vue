<template>
  <AppLayout>
    <TablePageLayout>
      <template #filters>
        <div class="flex flex-wrap items-center gap-3">
          <div class="flex-1 sm:max-w-72">
            <input
              v-model="searchQuery"
              type="text"
              :placeholder="t('admin.dispositions.searchPlaceholder')"
              class="input"
              @input="handleSearch"
            />
          </div>
          <Select
            v-model="statusFilter"
            :options="statusOptions"
            class="w-40"
            @change="handleStatusChange"
          />
          <div class="flex flex-1 justify-end">
            <button
              type="button"
              class="btn btn-secondary"
              :disabled="loading"
              :title="t('common.refresh')"
              @click="loadItems"
            >
              <Icon name="refresh" size="md" :class="loading ? 'animate-spin' : ''" />
            </button>
          </div>
        </div>
      </template>

      <template #table>
        <DataTable
          :columns="columns"
          :data="items"
          :loading="loading"
          default-sort-key="created_at"
          default-sort-order="desc"
        >
          <template #cell-user="{ row }">
            <div class="min-w-0">
              <div class="truncate font-medium text-gray-900 dark:text-white">{{ row.user.email }}</div>
              <div class="mt-1 flex items-center gap-2 text-xs text-gray-500 dark:text-dark-400">
                <span>#{{ row.user.id }}</span>
                <span v-if="row.user.username" class="truncate">{{ row.user.username }}</span>
              </div>
            </div>
          </template>

          <template #cell-status="{ row }">
            <span :class="['badge', row.is_disabled ? 'badge-danger' : 'badge-success']">
              {{ row.is_disabled ? t('admin.dispositions.statusDisabled') : t('admin.dispositions.statusActive') }}
            </span>
          </template>

          <template #cell-actions="{ row }">
            <div class="flex max-w-md flex-wrap gap-1.5">
              <span
                v-for="action in enabledActionLabels(row.actions)"
                :key="action"
                class="rounded-full bg-gray-100 px-2 py-0.5 text-xs text-gray-700 dark:bg-dark-700 dark:text-dark-200"
              >
                {{ action }}
              </span>
              <span v-if="enabledActionLabels(row.actions).length === 0" class="text-sm text-gray-400">
                {{ t('admin.dispositions.noActions') }}
              </span>
            </div>
          </template>

          <template #cell-reason="{ value }">
            <div class="max-w-md whitespace-pre-wrap text-sm text-gray-700 dark:text-dark-200">{{ value }}</div>
          </template>

          <template #cell-operator="{ row }">
            <span class="text-sm text-gray-500 dark:text-dark-400">
              {{ row.operator?.email || (row.operator?.id ? `#${row.operator.id}` : t('admin.dispositions.unknownOperator')) }}
            </span>
          </template>

          <template #cell-created_at="{ value }">
            <span class="text-sm text-gray-500 dark:text-dark-400">{{ formatDateTime(value) }}</span>
          </template>

          <template #cell-row_actions="{ row }">
	            <button
	              type="button"
	              class="btn btn-sm"
	              :class="canRestoreAccess(row) ? 'btn-primary' : 'btn-secondary'"
	              :disabled="!canRestoreAccess(row) || unbanningUserID === row.user.id"
	              @click="openUnban(row)"
	            >
	              {{ canRestoreAccess(row) ? t('admin.dispositions.unban') : t('admin.dispositions.alreadyActive') }}
	            </button>
          </template>

          <template #empty>
            <EmptyState
              :title="t('admin.dispositions.emptyTitle')"
              :description="t('admin.dispositions.emptyDescription')"
            />
          </template>
        </DataTable>
      </template>

      <template #pagination>
        <Pagination
          v-if="pagination.total > 0"
          :page="pagination.page"
          :total="pagination.total"
          :page-size="pagination.page_size"
          @update:page="handlePageChange"
          @update:pageSize="handlePageSizeChange"
        />
      </template>
    </TablePageLayout>

    <Teleport to="body">
      <div v-if="unbanTarget" class="fixed inset-0 z-[10000] flex items-center justify-center bg-black/40 p-4">
        <div class="w-full max-w-lg rounded-xl bg-white shadow-xl ring-1 ring-black/5 dark:bg-dark-800 dark:ring-white/10">
          <div class="border-b border-gray-100 px-5 py-4 dark:border-dark-700">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white">{{ t('admin.dispositions.unbanTitle') }}</h3>
            <p class="mt-1 text-sm text-gray-500 dark:text-dark-400">{{ unbanTarget.user.email }}</p>
          </div>
          <div class="space-y-3 px-5 py-4">
            <p class="text-sm text-gray-600 dark:text-dark-300">{{ t('admin.dispositions.unbanHint') }}</p>
            <div>
              <label class="input-label">{{ t('admin.dispositions.unbanReason') }}</label>
              <textarea
                v-model.trim="unbanReason"
                rows="3"
                class="input"
                :placeholder="t('admin.dispositions.unbanReasonPlaceholder')"
              />
            </div>
          </div>
          <div class="flex justify-end gap-3 border-t border-gray-100 px-5 py-4 dark:border-dark-700">
            <button type="button" class="btn btn-secondary" :disabled="Boolean(unbanningUserID)" @click="closeUnban">
              {{ t('common.cancel') }}
            </button>
            <button type="button" class="btn btn-primary" :disabled="Boolean(unbanningUserID)" @click="submitUnban">
              {{ unbanningUserID ? t('admin.dispositions.unbanning') : t('admin.dispositions.confirmUnban') }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { adminAPI } from '@/api/admin'
import type { UserDispositionAuditEntry } from '@/api/admin/users'
import { useAppStore } from '@/stores/app'
import { formatDateTime } from '@/utils/format'
import type { Column } from '@/components/common/types'
import AppLayout from '@/components/layout/AppLayout.vue'
import TablePageLayout from '@/components/layout/TablePageLayout.vue'
import DataTable from '@/components/common/DataTable.vue'
import Pagination from '@/components/common/Pagination.vue'
import EmptyState from '@/components/common/EmptyState.vue'
import Select from '@/components/common/Select.vue'
import Icon from '@/components/icons/Icon.vue'

const { t } = useI18n()
const appStore = useAppStore()

const loading = ref(false)
const items = ref<UserDispositionAuditEntry[]>([])
const searchQuery = ref('')
const statusFilter = ref<'active' | 'disabled' | ''>('disabled')
const unbanTarget = ref<UserDispositionAuditEntry | null>(null)
const unbanReason = ref('')
const unbanningUserID = ref<number | null>(null)
let searchTimer: number | undefined

const pagination = reactive({
  page: 1,
  page_size: 20,
  total: 0
})

const columns = computed<Column[]>(() => [
  { key: 'user', label: t('admin.dispositions.columns.user'), sortable: false },
  { key: 'status', label: t('admin.dispositions.columns.status'), sortable: false },
  { key: 'actions', label: t('admin.dispositions.columns.actions'), sortable: false },
  { key: 'reason', label: t('admin.dispositions.columns.reason'), sortable: false },
  { key: 'operator', label: t('admin.dispositions.columns.operator'), sortable: false },
  { key: 'created_at', label: t('admin.dispositions.columns.createdAt'), sortable: false },
  { key: 'row_actions', label: t('admin.dispositions.columns.manage'), sortable: false }
])

const statusOptions = computed(() => [
  { value: '', label: t('admin.dispositions.allAudited') },
  { value: 'disabled', label: t('admin.dispositions.onlyDisabled') },
  { value: 'active', label: t('admin.dispositions.onlyActive') }
])

const loadItems = async () => {
  loading.value = true
  try {
    const response = await adminAPI.users.listDispositions(pagination.page, pagination.page_size, {
      status: statusFilter.value,
      search: searchQuery.value.trim()
    })
    items.value = response.items
    pagination.total = response.total
    pagination.page = response.page
    pagination.page_size = response.page_size
  } catch (error: any) {
    appStore.showError(error.response?.data?.detail || error.message || t('admin.dispositions.failedToLoad'))
    console.error('Failed to load disposition users:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  if (searchTimer) window.clearTimeout(searchTimer)
  searchTimer = window.setTimeout(() => {
    pagination.page = 1
    loadItems()
  }, 300)
}

const handleStatusChange = () => {
  pagination.page = 1
  loadItems()
}

const handlePageChange = (page: number) => {
  pagination.page = page
  loadItems()
}

const handlePageSizeChange = (pageSize: number) => {
  pagination.page_size = pageSize
  pagination.page = 1
  loadItems()
}

const enabledActionLabels = (actions: Record<string, unknown>) => {
  return Object.entries(actions)
    .filter(([, value]) => value === true)
    .map(([key]) => actionLabel(key))
}

const hasDispositionDisabledAPIKeys = (row: UserDispositionAuditEntry) => {
  const ids = row.summary?.disabled_api_key_ids
  return row.actions?.disable_api_keys === true && Array.isArray(ids) && ids.length > 0
}

const canRestoreAccess = (row: UserDispositionAuditEntry) => {
  return row.is_disabled || hasDispositionDisabledAPIKeys(row)
}

const actionLabel = (key: string) => {
  const labels: Record<string, string> = {
    disable_user: t('admin.dispositions.actionLabels.disableUser'),
    disable_api_keys: t('admin.dispositions.actionLabels.disableApiKeys'),
    revoke_subscriptions: t('admin.dispositions.actionLabels.revokeSubscriptions'),
    clear_balance: t('admin.dispositions.actionLabels.clearBalance'),
    freeze_balance: t('admin.dispositions.actionLabels.freezeBalance'),
    append_note: t('admin.dispositions.actionLabels.appendNote'),
    unban_user: t('admin.dispositions.actionLabels.unbanUser'),
    enable_api_keys: t('admin.dispositions.actionLabels.enableApiKeys')
  }
  return labels[key] || key
}

const openUnban = (row: UserDispositionAuditEntry) => {
  unbanTarget.value = row
  unbanReason.value = ''
}

const closeUnban = () => {
  if (unbanningUserID.value) return
  unbanTarget.value = null
  unbanReason.value = ''
}

const submitUnban = async () => {
  if (!unbanTarget.value) return
  unbanningUserID.value = unbanTarget.value.user.id
  try {
    const result = await adminAPI.users.unbanDispositionUser(
      unbanTarget.value.user.id,
      unbanReason.value.trim() || t('admin.dispositions.defaultUnbanReason')
    )
    appStore.showSuccess(t('admin.dispositions.unbanSuccess', { keys: result.enabled_api_keys }))
    unbanTarget.value = null
    unbanReason.value = ''
    loadItems()
  } catch (error: any) {
    appStore.showError(error.response?.data?.detail || error.message || t('admin.dispositions.unbanFailed'))
    console.error('Failed to unban user:', error)
  } finally {
    unbanningUserID.value = null
  }
}

onMounted(loadItems)
</script>
