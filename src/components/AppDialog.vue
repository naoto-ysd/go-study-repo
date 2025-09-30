<template>
  <dialog ref="el" @cancel.prevent="close" @click.self="close" :aria-label="label" aria-modal="true">
    <slot />
  </dialog>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue';

const props = defineProps({
  show: { type: Boolean, required: true },
  label: { type: String, default: 'ダイアログ' }
});
const emit = defineEmits(['update:show','close']);

const el = ref(null);
const close = () => {
  emit('update:show', false);
  emit('close');
};

watch(() => props.show, (val) => {
  if (!el.value) return;
  if (val) {
    if (!el.value.open) el.value.showModal();
  } else if (el.value.open) {
    el.value.close();
  }
});

onMounted(() => { if (props.show && el.value && !el.value.open) el.value.showModal(); });
</script>

<style scoped>
:host, dialog {
  font: inherit;
}
</style>
