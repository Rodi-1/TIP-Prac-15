# ПЗ №15 — Модульные тесты Go

Сортов Семён Александрович, ЭФМО-01-26. Дата: 25.10.2026.

Три пакета: mathx (Sum, Divide, MustDivide), stringsx (Clip) и service (FindIDByEmail через UserRepo). Тесты используют testing, таблицы, t.Run, testify/assert и require. Сервис изолирован от БД локальным stubRepo, который также проверяет email и число обращений. Требуется Go 1.25+, внешние сервисы не нужны.

```powershell
go test ./...
go test -v ./internal/...
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

Открыть coverage.html в браузере. В целевых пакетах покрытие 100%; проверяются результат, ветвления и ошибки, а не только сам факт вызова. Есть положительные/отрицательные числа, ноль, целочисленное деление, деление на ноль с ErrorIs, пустая строка и границы max, успех/отсутствие/ошибка репозитория. Clip работает по байтам, как допускает задание; его тесты не предполагают сохранение целых UTF-8 символов.

## Все бонусные пункты

```powershell
go test -bench . -benchmem ./internal/mathx
go test -run '^$' -fuzz FuzzClip -fuzztime 2s -parallel 4 ./internal/stringsx
```

BenchmarkSum записывает результат в переменную, чтобы он не был полностью отброшен. TestMustDivide проверяет require.Panics. FuzzClip проверяет свойства длины и префикса, включая случайные входы. Результаты бенчмарка зависят от оптимизации компилятора; у столь простой функции значителен вклад самого цикла.
