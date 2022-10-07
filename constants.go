package main

const (
	POLLING_INTERVAL_IN_MS             = 30000 // 30 secs
	BROWSERSTACK_DOMAIN                = "https://api-cloud.browserstack.com"
	APP_UPLOAD_ENDPOINT                = "/app-automate/xcuitest/v2/app"
	TEST_SUITE_UPLOAD_ENDPOINT         = "/app-automate/xcuitest/v2/test-suite"
	APP_AUTOMATE_BUILD_ENDPOINT        = "/app-automate/xcuitest/v2/build"
	APP_AUTOMATE_BUILD_STATUS_ENDPOINT = "/app-automate/xcuitest/v2/builds/"
	APP_AUTOMATE_BUILD_DASHBOARD_URL   = "https://app-automate.browserstack.com/dashboard/v2/builds/"
	TEST_RUNNER_RELATIVE_PATH_BITRISE  = "/Debug-iphoneos/Tests iOS-Runner.app"
	TEST_RUNNER_ZIP_FILE_NAME          = "test_suite.zip"

	SAMPLE_APP        = "bs://b91841adbf33515fef7a1cca869a9526a86f9a0e"
	SAMPLE_TEST_SUITE = "bs://535a0932c8a785384b8470ec6166e093cd3b2c5f"
	SAMPLE_BUILD_ID   = "56fec97937b22c785a6c5e08c13f629d505f5cd9"

	UPLOAD_APP_ERROR         = "Failed to upload app on BrowserStack, error : %s"
	FILE_NOT_AVAILABLE_ERROR = "Failed to upload test suite on BrowserStack, error: file not available"
	INVALID_FILE_TYPE_ERROR  = "Failed to upload test suite on BrowserStack, error: invalid file type"
	BUILD_FAILED_ERROR       = "Failed to execute build on BrowserStack, error: %s"
	FETCH_BUILD_STATUS_ERROR = "Failed to fetch test results, error: %s"
	HTTP_ERROR               = "Something went wrong while processing your request, error: %s"
	RUNNER_APP_NOT_FOUND     = "xcuitest_testsuite_path: couldn’t find the <AppnameUITests>-Runner.app .  Please add the $BITRISE_TEST_BUNDLE_PATH from Xcode Build for testing for iOS step or the absolute path of <AppnameUITests>-Runner.app"
	IPA_NOT_FOUND            = "Failed to generate an .ipa file. Please verify the value in $BUNDLE_APP_NAME"
	FILE_NOT_FOUND           = "File not found: %s"
	FILE_COPY_ERROR          = "Failed to copy file, error: %s"
	FILE_DIR_ERROR           = "Failed to create directory, error: %s"
	FILE_ZIP_ERROR           = "Failed to zip file, error: %s"
)
