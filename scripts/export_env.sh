#!/usr/bin/env sh

ENV_FILE="$1"

if [ -z "$ENV_FILE" ]; then
  echo "Usage: source $0 path/to/.env" >&2
  return 1 2>/dev/null || exit 1
fi

if [ ! -f "$ENV_FILE" ]; then
  echo "File not found: $ENV_FILE" >&2
  return 1 2>/dev/null || exit 1
fi

while IFS='=' read -r key val; do
  case "$key" in
    ''|\#*) continue ;;
  esac

  val=$(echo "$val" | sed -e 's/^["'\'']//;s/["'\'']$//')

  key=$(echo "$key" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')

  export "$key=$val"
done < "$ENV_FILE"

