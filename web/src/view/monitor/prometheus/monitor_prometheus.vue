<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="名称">
          <el-input
            v-model="searchInfo.name"
            placeholder="名称"
          />
        </el-form-item>
        <el-form-item label="编码">
          <el-input
            v-model="searchInfo.code"
            placeholder="编码"
          />
        </el-form-item>
        <el-form-item label="url">
          <el-input
              v-model="searchInfo.url"
              placeholder="url"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >
            查询
          </el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>



    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
            type="primary"
            icon="plus"
            @click="openDialog('addApi')"
        >
          新增
        </el-button>
        <el-button
            icon="delete"
            :disabled="!apis.length"
            @click="onDelete"
        >
          删除
        </el-button>
      </div>
      <el-table
        :data="tableData"
        @sort-change="sortChange"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="id"
          width="100"
          prop="ID"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="名称"
          width="200"
          prop="name"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="编码"
          width="200"
          prop="code"
          sortable="custom"
        />
        <el-table-column label="地址" prop="url">
          <template v-slot="scope">
            <a :href="scope.row.url" target="_blank" class="buttonText" style="text-decoration: none;color:#409eff">{{scope.row.url}}</a>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          fixed="right"
          label="操作"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editPrometheusFunc(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              icon="delete"

              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer
        v-model="dialogFormVisible"
        size="60%"
        :before-close="closeDialog"
        :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ dialogTitle }}</span>
          <div>
            <el-button @click="closeDialog">
              取 消
            </el-button>
            <el-button
                type="primary"
                @click="enterDialog"
            >
              确 定
            </el-button>
          </div>
        </div>
      </template>

      <warning-bar title="新增Prometheus，需要在角色管理内配置权限才可使用" />
      <el-form
          ref="prometheusForm"
          :model="form"
          :rules="rules"
          label-width="80px"
      >
        <el-form-item
            label="名称"
            prop="name"
        >
          <el-input
              v-model="form.name"
              autocomplete="off"
          />
        </el-form-item>
        <el-form-item
            label="编码"
            prop="code"
        >
          <el-input
              v-model="form.code"
              autocomplete="off"
          />
        </el-form-item>
        <el-form-item
            label="url"
            prop="url"
        >
          <el-input
              v-model="form.url"
              autocomplete="off"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  getPrometheusList,
  getPrometheusById
} from '@/api/monitor_prometheus'
import { toSQLLine } from '@/utils/stringFun'
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'


defineOptions({
  name: 'Prometheus',
})

const methodFilter = (value) => {
  const target = methodOptions.value.filter(item => item.value === value)[0]
  return target && `${target.label}`
}

const apis = ref([])
const form = ref({
  url: '',
  name: '',
  code: ''
})
const methodOptions = ref([
  {
    value: 'POST',
    label: '创建',
    type: 'success'
  },
  {
    value: 'GET',
    label: '查看',
    type: ''
  },
  {
    value: 'PUT',
    label: '更新',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: '删除',
    type: 'danger'
  }
])

const type = ref('')
const rules = ref({
  url: [{ required: true, message: '请输入url', trigger: 'blur' }],
  name: [
    { required: true, message: '请输入名称', trigger: 'blur' }
  ],
  code: [
    { required: true, message: '请选择编码', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})


const onReset = () => {
  searchInfo.value = {}
}
// 搜索

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getPrometheusList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
// 批量操作
const handleSelectionChange = (val) => {
  apis.value = val
}

const editPrometheusFunc = async(row) => {
  const res = await getPrometheusById({ id: row.ID })
  form.value = res.data.prometheus
  openDialog('edit')
}

// 弹窗相关
const prometheusForm = ref(null)
const initForm = () => {
  prometheusForm.value.resetFields()
  form.value = {
    name: '',
    code: '',
    url: ''
  }
}

const dialogTitle = ref('新增Prometheus')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = '新增Prometheus'
      break
    case 'edit':
      dialogTitle.value = '编辑Prometheus'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
</style>
