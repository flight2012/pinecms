<template>
	<div class="cl-menu-tree">
		<el-popover
			placement="bottom-start"
			trigger="click"
			width="500px"
			popper-class="popper-menu-tree"
		>
			<el-input v-model="keyword" size="small">
				<template #prefix>
					<i class="el-input__icon el-icon-search"></i>
				</template>
			</el-input>

			<el-tree
				ref="treeRef"
				node-key="menuId"
				:data="treeList"
				:props="{
					label: 'name',
					children: 'children'
				}"
				:highlight-current="true"
				:expand-on-click-node="false"
				:default-expanded-keys="expandedKeys"
				:filter-node-method="filterNode"
				@current-change="onCurrentChange"
			/>

			<template #reference>
				<el-input v-model="name" readonly placeholder="请选择" />
			</template>
		</el-popover>
	</div>
</template>

<script lang="ts">
import { computed, defineComponent, inject, onMounted, ref, watch } from "vue";
import { deepTree } from "/@/core/utils";

export default defineComponent({
	name: "el-category-tree",

	props: {
		modelValue: [Number, String]
	},

	emits: ["update:modelValue"],

	setup(props, { emit }) {
		// 请求服务
		const service = inject<any>("service");

		// 关键字
		const keyword = ref<string>("");

		// 树形列表
		const list = ref<any[]>([]);

		// 展开值
		const expandedKeys = ref<any[]>([]);

		// el-tree 组件
		const treeRef = ref<any>({});

		// 绑定值回调
		function onCurrentChange({ id }: any) {
			emit("update:modelValue", id);
		}

		// 过滤节点
		function filterNode(value: string, data: any) {
			if (!value) return true;
			return data.name.indexOf(value) !== -1;
		}

		// 节点名称
		const name = computed(() => {
			const item = list.value.find((e) => e.id == props.modelValue);
			return item ? item.name : "一级栏目";
		});

		// 树形列表
		const treeList = computed(() => deepTree(list.value));

		// 监听关键字过滤
		watch(keyword, (val: string) => {
			treeRef.value.filter(val);
		});

		onMounted(async function () {
			const ret = await service.system.category.list();
			const _list = ret.list.filter((e: any) => e.type != 2);
			_list.unshift({
				name: "一级栏目",
				id: null
			});
			list.value = _list;
		});

		return {
			keyword,
			list,
			expandedKeys,
			treeRef,
			name,
			treeList,
			filterNode,
			onCurrentChange
		};
	}
});
</script>

<style lang="scss">
.popper-menu-tree {
	box-sizing: border-box;

	.el-input {
		margin-bottom: 10px;
	}
}
</style>
