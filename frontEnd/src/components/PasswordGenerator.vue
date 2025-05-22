<template>
  <form @submit.prevent="submit">
    <div>
      <label for="length">Password length:</label>
      <input id="length" v-model.number="length" type="number" />
    </div>

    <div v-if="options.length">
      <div v-for="option in options" :key="option.key">
        <label>
          <input
            type="checkbox"
            :value="option.key"
            v-model="selectedOptions"
          />
          {{ option.description }}
        </label>
      </div>
    </div>

    <button type="submit">Generate</button>

    <div v-if="generatedPassword" class="result">
      <p><strong>Generated password:</strong> {{ generatedPassword }}</p>
    </div>

    <div v-if="errors.length" class="errors">
      <p v-for="(err, idx) in errors" :key="idx" class="error">{{ err }}</p>
    </div>
  </form>
</template>

<script setup>
import { ref, onMounted } from "vue";

const length = ref(8);
const options = ref([]);
const selectedOptions = ref([]);
const generatedPassword = ref("");
const errors = ref([]);

onMounted(async () => {
  const res = await fetch(
    "http://localhost:8080/api/password/generationOptions"
  );
  const result = await res.json();
  options.value = result.data?.generationOptions || [];
});

async function submit() {
  errors.value = [];
  generatedPassword.value = "";

  const payload = {
    length: length.value,
    options: selectedOptions.value,
  };

  const res = await fetch("http://localhost:8080/api/password/generate", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  const result = await res.json();

  if (res.ok && result.isSuccess) {
    generatedPassword.value = result.data.password;
  } else {
    errors.value = result.data?.errors || [
      result.message || "Something went wrong.",
    ];
  }
}
</script>

<style scoped>
form {
  max-width: 400px;
  margin: 2rem auto;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.result {
  margin-top: 1rem;
  padding: 1rem;
  background: #f0f0f0;
  border-radius: 6px;
  font-family: monospace;
}

.errors {
  color: red;
}

.error {
  margin: 0.25rem 0;
}
</style>
