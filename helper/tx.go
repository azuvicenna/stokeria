package helper

import "gorm.io/gorm"

// CommitOrRollback dipanggil menggunakan defer pada layer Service.
// Fungsi ini menangkap panic dari bawah, me-rollback transaksi, lalu melempar panic kembali ke atas.
func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		PanicIfError(errorRollback)
		panic(err) // Lempar kembali panic agar ditangkap oleh Global Error Handler Fiber
	} else {
		errorCommit := tx.Commit().Error
		PanicIfError(errorCommit)
	}
}