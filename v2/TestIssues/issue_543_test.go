package TestIssues

import (
	"testing"
)

func TestIssue543(t *testing.T) {
	expected := "CASE  WHEN \"SESSION_ID\" IS NULL THEN \"SECURITY_GROUP_ID\" END "
	db, err := getDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err = db.Close()
		if err != nil {
			t.Error(err)
		}
	}()
	var result string
	err = db.QueryRow("select column_expression from dba_ind_expressions where index_name = 'members_last_name_fi' and column_position = 1").Scan(&result)
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("expected: %s and got: %s", expected, result)
	}
}
