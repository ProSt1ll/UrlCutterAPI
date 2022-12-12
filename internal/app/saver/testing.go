package saver

//func TestSaver(t *testing.T, dbname string) (*DBSaver, func(...string)) {
//t.Helper()
//db := NewDBSaver("localhost", "5436", "test_db")
//
//if err := db.db.Open(); err != nil {
//	t.Fatal(err)
//}
//
//return db, func(tables ...string) {
//	if len(tables) > 0 {
//		_, err := db.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
//		if err != nil {
//			t.Fatal(err)
//		}
//		_, err = db.db.Exec(fmt.Sprintf("ALTER SEQUENCE urls_id_seq RESTART WITH 1"))
//		if err != nil {
//			t.Fatal(err)
//		}
//
//	}
//	if err := db.Close(); err != nil {
//		t.Fatal(err)
//	}
//}
//}
