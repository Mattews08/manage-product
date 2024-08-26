export const formatDate = (dateStr: string) => {
  if (!dateStr) {
    return "Data inválida";
  }

  const date = new Date(dateStr);

  if (isNaN(date.getTime())) {
    return "Data inválida";
  }

  return `${date.getDate().toString().padStart(2, "0")}-${(date.getMonth() + 1)
    .toString()
    .padStart(2, "0")}-${date.getFullYear()}`;
};
