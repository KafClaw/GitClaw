#!/usr/bin/env bash
set -euo pipefail

# Compatibility wrapper: delegate to canonical script.
script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
canonical_script="${script_dir}/../../gitea/scripts/enroll.sh"

if [[ ! -x "${canonical_script}" ]]; then
  echo "canonical enroll script not found or not executable: ${canonical_script}" >&2
  exit 1
fi

exec "${canonical_script}" "$@"
