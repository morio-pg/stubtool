import moment from "moment";

export function dateFormat(date) {
  return moment(new Date(date)).format("YYYY/MM/DD HH:mm");
}

export function truncate(value, maxLength) {
  if (!value) {
    return "-";
  } else if (value.length <= maxLength) {
    return value;
  } else {
    return value.substring(0, maxLength - 1) + "…";
  }
}
