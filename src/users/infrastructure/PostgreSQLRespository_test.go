package infrastructure_test

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/luizarnoldch/mesapara2_backend/config/env"
// )

// // Unable to open connection with PostgreSQL
// func TestUnableToOpenConnectionWithPostgreSQL(t *testing.T) {
//     // Mock the LoadConfig function to return an error
//     envLoadConfig = func(filePath string) (*env.CONFIG, error) {
//         return nil, fmt.Errorf("error loading .env")
//     }

//     // Call the GetPostgreSQLClient function
//     client := GetPostgreSQLClient()

//     // Assert that the client is nil
//     if client != nil {
//         t.Errorf("Expected client to be nil, but got %v", client)
//     }
// }

// // Unable to load configuration file
// func TestGetPostgreSQLClientUnableToLoadConfigFile(t *testing.T) {
//     // Create a temporary directory for the .env file
//     tempDir, err := ioutil.TempDir("", "test")
//     if err != nil {
//         t.Fatalf("Failed to create temporary directory: %v", err)
//     }
//     defer os.RemoveAll(tempDir)

//     // Set the file path to a non-existent file
//     filePath := filepath.Join(tempDir, "nonexistent.env")

//     // Call the GetPostgreSQLClient function
//     client := GetPostgreSQLClient()

//     // Assert that the client is nil
//     if client != nil {
//         t.Errorf("Expected client to be nil, but got %v", client)
//     }
// }


// // Verify correct dataSource string is generated
// func TestVerifyCorrectDataSourceStringGenerated(t *testing.T) {
//     // Set up test data
//     expectedDataSource := "host=localhost port=5432 user=test password=test dbname=test sslmode=disable"

//     // Call the function
//     client := GetPostgreSQLClient()

//     // Check the result
//     if client.DataSource() != expectedDataSource {
//         t.Errorf("Incorrect dataSource string generated. Expected: %s, Got: %s", expectedDataSource, client.DataSource())
//     }
// }


// // Set connection max lifetime, max open connections and max idle connections
// func test_set_connection_parameters(t *testing.T) {
//     // Set up
//     config := &CONFIG{
//         MICRO: MICRO{
//             DB: DB{
//                 PSQL: PSQL{
//                     USER:   "test_user",
//                     PASS:   "test_password",
//                     HOST:   "localhost",
//                     PORT:   "5432",
//                     SCHEMA: "test_schema",
//                 },
//             },
//         },
//     }

//     // Execute
//     client := GetPostgreSQLClient()

//     // Assert
//     if client.MaxConnLifetime() != time.Minute*3 {
//         t.Errorf("Expected connection max lifetime to be %v, but got %v", time.Minute*3, client.MaxConnLifetime())
//     }
//     if client.MaxOpenConns() != 10 {
//         t.Errorf("Expected max open connections to be %d, but got %d", 10, client.MaxOpenConns())
//     }
//     if client.MaxIdleConns() != 10 {
//         t.Errorf("Expected max idle connections to be %d, but got %d", 10, client.MaxIdleConns())
//     }
// }


// // Unable to ping PostgreSQL connection
// func TestGetPostgreSQLClient_UnableToPing(t *testing.T) {
//     // Set up test environment
//     os.Setenv("API_HOST", "localhost")
//     os.Setenv("API_PORT", "5432")
//     os.Setenv("DB_USER", "testuser")
//     os.Setenv("DB_PASSWORD", "testpassword")
//     os.Setenv("DB_HOST", "localhost")
//     os.Setenv("DB_PORT", "5432")
//     os.Setenv("DB_SCHEMA", "testdb")

//     // Mock the LoadConfig function to return an error
//     env.LoadConfig = func(filePath string) (*env.CONFIG, error) {
//         return nil, fmt.Errorf("error loading .env file")
//     }

//     // Call the function under test
//     client := GetPostgreSQLClient()

//     // Assert that the client is nil
//     if client != nil {
//         t.Errorf("Expected client to be nil, but got %v", client)
//     }
// }


// // Verify connection is closed after function execution
// func TestVerifyConnectionClosed(t *testing.T) {
//     // Set up
//     mockDB := &sqlx.DB{}
//     mockDBClosed := false
//     mockDB.CloseFunc = func() error {
//         mockDBClosed = true
//         return nil
//     }
//     sqlxOpenFunc := func(driverName, dataSourceName string) (*sqlx.DB, error) {
//         return mockDB, nil
//     }
//     sqlxOpen = sqlxOpenFunc

//     // Execute the function
//     GetPostgreSQLClient()

//     // Verify connection is closed
//     if !mockDBClosed {
//         t.Errorf("Expected connection to be closed, but it was not")
//     }
// }


// // Verify correct SQL driver is used
// func TestVerifyCorrectSQLDriverIsUsed(t *testing.T) {
//     // Set up the test environment
//     os.Setenv("API_HOST", "localhost")
//     os.Setenv("API_PORT", "5432")
//     os.Setenv("DB_USER", "testuser")
//     os.Setenv("DB_PASSWORD", "testpassword")
//     os.Setenv("DB_HOST", "localhost")
//     os.Setenv("DB_PORT", "5432")
//     os.Setenv("DB_SCHEMA", "testdb")

//     // Call the function under test
//     client := GetPostgreSQLClient()

//     // Assert that the correct SQL driver is used
//     driverName := client.DriverName()
//     expectedDriverName := "postgres"
//     if driverName != expectedDriverName {
//         t.Errorf("Expected SQL driver name to be %s, but got %s", expectedDriverName, driverName)
//     }
// }


// // Verify correct values are extracted from configuration file
// func TestGetPostgreSQLClient(t *testing.T) {
//     // Create a temporary .env file with test values
//     tempFile, err := ioutil.TempFile("", ".env")
//     if err != nil {
//         t.Fatalf("failed to create temporary .env file: %v", err)
//     }
//     defer os.Remove(tempFile.Name())

//     // Write test values to the temporary .env file
//     _, err = tempFile.WriteString("API_HOST=test_host\nAPI_PORT=test_port\nDB_USER=test_user\nDB_PASS=test_password\nDB_HOST=test_db_host\nDB_PORT=test_db_port\nDB_SCHEMA=test_db_schema")
//     if err != nil {
//         t.Fatalf("failed to write to temporary .env file: %v", err)
//     }

//     // Close the temporary .env file
//     err = tempFile.Close()
//     if err != nil {
//         t.Fatalf("failed to close temporary .env file: %v", err)
//     }

//     // Call the GetPostgreSQLClient function
//     client := GetPostgreSQLClient()

//     // Verify that the client is not nil
//     if client == nil {
//         t.Error("expected client to not be nil")
//     }

//     // Verify that the database connection parameters are correct
//     expectedDataSource := "host=test_db_host port=test_db_port user=test_user password=test_password dbname=test_db_schema sslmode=disable"
//     if client.DriverName() != "postgres" || client.DataSourceName() != expectedDataSource {
//         t.Errorf("unexpected database connection parameters, got: %s, want: %s", client.DataSourceName(), expectedDataSource)
//     }
// }


