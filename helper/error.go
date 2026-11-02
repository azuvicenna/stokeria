package helper

// PanicIfError akan menghentikan eksekusi secara paksa jika err != nil.
// Fungsi ini hanya boleh digunakan pada proses inisialisasi (layer app)
// atau di dalam handler yang sudah dilindungi oleh fitur Recover milik Fiber.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}