# 🚀 Шпаргалка по работе с Git в отдельной ветке

## 1. Схлопывание коммитов (Squash Commit) — Локально
*Делается строго **ДО** отправки ветки на сервер, чтобы очистить историю.*
* **Легкий откат:** Если найдут баг, всю фичу можно отменить командой `git revert` за один клик.

### Как сделать:
1. Запустите интерактивный rebase для всех коммитов вашей ветки относительно main:
   ```bash
   git rebase -i main
   ```
2. В текстовом редакторе замените слово `pick` на **`squash`** (или **`s`**) у всех коммитов, кроме **первого**:
   ```text
   pick a1b2c3d Реализована фича Х (основной)
   squash e5f6g7h исправил опечатку
   squash i9j0k1l забыл точку с запятой
   ```

⚠️ Если вы уже отправляли эту ветку на GitHub до схлопывания, `push` не пройдет. Отправить изменения принудительно: `git push --force-with-lease`.

---

## 2. Отправка ветки на сервер
После того как коммиты объединены, отправьте ветку в удаленный репозиторий (eсли вы планируете делать Squash на GitHub, локальный интерактивный rebase (шаг 1) не нужен):
```bash
git push -u origin имя-ветки
```
*Флаг `-u` свяжет ветки, и дальше в этой ветке можно будет писать `git push`.*

---

## 3. Слияние с основной веткой (main)

1. **Находясь в своей ветке обновите локальный main**:
   ```bash
   git fetch origin
   ```
2. **Накатите рабочую ветку на актуальный main**:
   ```bash
   git rebase origin/main
   ```
   *Рабочая ветка растет прямо из самой свежей точки `main`.*
3. **Влейте изменения в main через Fast-Forward**:
   ```bash
   git switch main
   git merge имя-ветки
   ```

* **`git rebase main` (в своей ветке):** Git **ставит рабочую ветку на свежий фундамент `main`**.
* Рабочая ветка уже стоит на вершине `main`, `git merge` просто сдвигает указатель `main` вперед (**Fast-forward перемотка**).
* **⚠️ НЕЛЬЗЯ `git rebase имя-ветки` на ветке `main`?** `main` перепишется поверх временной ветки.

1. **Отправьте обновленный main на сервер**:
   ```bash
   git push origin main
   ```

---

## 4. Удаление ветки

* **Локально:**
  ```bash
  git switch main
  git pull
  git branch -d имя-ветки
  # Allias deletes local branches that no longer exist on GitHub
  git gone

  git push origin --delete имя-ветки # На сервере
  ```

When PR is merged and you are ready to publish a new stable version:

```bash
git co main
git pull
# Trigger the semantic versioning and CHANGELOG.md manager
release
git push --follow-tags origin main
```
