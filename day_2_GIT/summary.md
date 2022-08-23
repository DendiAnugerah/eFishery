# Summary Day 2

Pada hari kedua ini, kita belajar mengenai pengenalan dan juga fundamental daripada Git,\
kita juga belajar mengenai beberapa fitur _command_ git yang sering digunakan,\
yang mana diantaranya adalah:

## Git Command

| Command                            | Kegunaan                                                             |
| ---------------------------------- | -------------------------------------------------------------------- |
| `git init <options> `              | Membuat repository pada file lokal                                   |
| `git add <file> / <directory>`     | Menambah file baru pada repository yang telah kita pilih             |
| `git remote <options> `            | Menambahkan koneksi remote                                           |
| `git status `                      | Untuk mengetahui status repository tersebut                          |
| `git commit <options> `            | Menyimpan perubahan yang sudah kita lakukan                          |
| `git push <options> `              | Mengupload perubahan file yang sudah kita commit sebelumnya          |
| `git fetch <remote> [options] `    | Mengambil commit terbaru                                             |
| `git pull <remote> <branch-name> ` | Mengambil commit terbaru lalu menggabungkan dengan branch aktif kita |
| `git branch <branch> `             | Melihat branch yang ada pada repository kita                         |
| `git merge <branch-name> `         | Menggabungkan dua branch menjadi satu                                |
| `git stash <command> `             | Untuk menyimpan proggres yang sudah kita lakukan                     |

## Commit Convention

Selain daripada itu, kita juga belajar mengenai Commit Convention menggunakan [Git Karma convention](karma-runner.github.io/),
yang mana ternyata dalam melakukan penamaan sebuah commit itu tidak bisa sembarangan, salah satunya agar dalam melakukan navigasi melalui history gitnya itu dapat dilakukan dengan mudah.\
Allowed `<type>`values:

- **feat** for a new feature for the user, not a new feature for build script. Such commit will trigger a release bumping a MINOR version.
- **fix** for a bug fix for the user, not a fix to a build script. Such commit will trigger a release bumping a PATCH version.
- **perf** for performance improvements. Such commit will trigger a release bumping a PATCH version.
- **docs** for changes to the documentation.
- **style** for formatting changes, missing semicolons, etc.
- **refactor** for refactoring production code, e.g. renaming a variable.
- **test** for adding missing tests, refactoring tests; no production code change.
- **build** for updating build configuration, development tools or other changes irrelevant to the user.

## Semantic Versioning

Dalam penamaan versi, ternyata juga memiliki aturan atau standar baku agar lebih memudahkan pemberian versi.

Contoh: v1.2.0\
`v<major>.<minor>.<patch>`

**Major**: mengubah API\
**Minor**: feat\
**Patch**: fix, perf

## Git Management

Diakhir sesi, dikenalkan juga [Trunk Base Development](https://trunkbaseddevelopment.com/), yaitu Git Management yang pada saat ini sedang digunakan oleh **eFishery**.

Trunk Base Development ini pada intinya merupakan sebuah model dalam pengontrolan/manajemen branch, agar memudahkan para developer untuk dapat berkolaborasi pada single branch, adanya Trunk Base Development ini salah satunya alasannya adalah, pada saat ini git workflow sudah kurang relevan jika digunakan dalam jumlah yang banyak.
