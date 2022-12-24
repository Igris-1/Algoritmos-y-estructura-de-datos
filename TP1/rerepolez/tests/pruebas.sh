#!/usr/bin/env bash

set -eu

PROGRAMA="$1"

RET=0
OUT=`mktemp`
trap "rm -f $OUT" EXIT

for x in *.test; do
  b=${x%.test}
  echo -n "Prueba $b... "
  cat "$x"

  ($PROGRAMA ${b}_partidos ${b}_padron <${b}_in || RET=$?) |
    diff -u --label "${b}_cátedra" --label "${b}_estudiante" ${b}_out - >$OUT || :

  if [[ $RET -ne 0 ]]; then
    echo -e "programa abortó con código $RET."
    exit $RET

  elif [[ -s $OUT ]]; then
    echo -e "output incorrecto:\n"
    cat $OUT
    exit 1

  else
    echo -e "OK."
  fi
  echo
done

echo "Prueba sin parámetros"
($PROGRAMA < 01_in || RET=$?) |
  diff -u --label "sin_parametros_cátedra" --label "sin_parametros_estudiante" sin_params_out - >$OUT || :

  if [[ $RET -ne 0 ]]; then
      echo -e "programa abortó con código $RET."
      exit $RET

  elif [[ -s $OUT ]]; then
      echo -e "output incorrecto:\n"
      cat $OUT
      exit 1

  else
      echo -e "OK."
  fi
  echo

echo "Prueba falando parámetros"
($PROGRAMA 01_partidos < 01_in || RET=$?) |
  diff -u --label "falta_parametro_cátedra" --label "falta_parametro_estudiante" sin_params_out - >$OUT || :

  if [[ $RET -ne 0 ]]; then
      echo -e "programa abortó con código $RET."
      exit $RET

  elif [[ -s $OUT ]]; then
      echo -e "output incorrecto:\n"
      cat $OUT
      exit 1

  else
      echo -e "OK."
  fi
