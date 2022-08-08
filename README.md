# BrowserStack App Automate - XCUI

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/en/steps-and-workflows/introduction-to-workflows.html).

OR

You can also run this step directly with [Bitrise CLI]([https://github.com/bitrise-io/bitrise](https://github.com/bitrise-io/bitrise#bitrise-offline-cli)).

## Run XCUI tests on BrowserStack

<details>
<summary>Description</summary>

Run your XCUI tests on BrowserStack App Automate. This step collects the built IPA from `$BITRISE_IPA_PATH` and the output bundle file from `$BITRISE_TEST_BUNDLE_PATH` environment variables.

## Configure the Step

Complete the following steps to configure BrowserStack's XCUI step in Bitrise:

1. Open the Workflow you want to use in the Workflow Editor.
​
2. Add the [Xcode Archive & Export for iOS](https://www.bitrise.io/integrations/steps/xcode-archive) and [Xcode Build for testing for iOS](https://www.bitrise.io/integrations/steps/xcode-build-for-test) steps to your workflow and configure them.
​
3. Add the **BrowserStack App Automate - XCUI** step below the **Xcode Archive & Export for iOS** and **Xcode Build for testing for iOS** steps.
​
4. Add your BrowserStack Username and Access Key in the **Authentication** step input.
​
5. For the **iOS app under test** input, the **BITRISE_IPA_PATH** output variable from the **Xcode Archive & Export for iOS** step exports the IPA file. Add `$BITRISE_IPA_PATH` to the **iOS app under test** input.<br /><br /> For the **XCUI test suite** input, the **BITRISE_TEST_BUNDLE_PATH** output variable from the **Xcode Build for testing for iOS step** exports the test suite. Add `$BITRISE_TEST_BUNDLE_PATH` to the **iOS app under test** input.<br /><br /> If you are not using  **Xcode Archive & Export for iOS** and **Xcode Build for testing for iOS** steps, ensure that the **iOS app under test** input points to the path of your app (`.ipa` file). Also, ensure that the **XCUI test suite** input points to the test suite runner file. In the case of the runner app, it should be in the `<any_path>/Debug-iphoneos` directory if you are providing an absolute path.<br />
​
6. Add one or more devices in the **Devices** step input.
​
7. Configure additional step inputs like **Debug logs** and **Test Configurations** and start your build.

</details>

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `iOS app` | Set the path of the app (.ipa) file. | Required | N/A |
| `XCUI test suite` | Set the path of the output bundle file. | Required | N/A |
| `Devices` | Provide one or more device-OS combination in a new line. For example: <br /> `iPhone 11-13` <br />`iPhone XS-15` | Required | N/A |
| `Instrumentation logs` | Generate instrumentation logs of the test session  | Optional | `true` |
| `Network logs` | Generate network logs of your test sessions to capture network traffic, latency, etc. | Optional | `false` |
| `Device Logs` | Generate device logs | Optional | `false` |
| `Capture screenshots` | Capture the screenshots of the test execution| Optional | `false` |
| `Video recording` | Record video of the test execution  | Optional | `true` |
| `Project name` | Project name of the tests | Optional | N/A |
| `Notify project status` | A callback URL to enable BrowserStack notify about completion of build under a given project.   | Optional | N/A  |
| `Local testing` | Enable local testing to retrieve app data hosted on local/private servers  |  Optional | `false` |
| `Test sharding` | Enable test sharding to split tests cases into different groups instead of running them sequentially. <br />Add the sharding value json here. <br />Examples: <br />**Input for only-testing strategy**: <br /> ```{"numberOfShards": 2, "mapping": [{"name": "Shard 1", "strategy": "only-testing", "values": ["SampleXCUITestsClass/testAlert", "SampleXCUITestsClass/testText"]}, {"name": "Shard 2", "strategy": "only-testing", "values": ["SampleXCUITestsClass/testLogin"]}]}``` <br /> **Input for skip-testing strategy**: ```{"numberOfShards": 2, "mapping": [{"name": "Shard 1", "strategy": "skip-testing", "values": ["SampleXCUITestsClass/testAlert"]}, {"name": "Shard 2", "strategy": "skip-testing", "values": ["SampleXCUITestsClass/testText"]}]}```|  Optional | N/A |
| `Filter test cases` | Provide comma-separated list of classes followed by the supported filtering strategy name `only-testing` and `skip-testing`.  <br /> Examples: <br />**For only-testing filtering strategy**: `only-testing SampleXCUITestsClass/testAlert, only-testing SampleXCUITestsClass/testText` <br /> **For skip-testing filtering strategy**: `skip-testing SampleXCUITestsClass/testAlert, skip-testing SampleXCUITestsClass/testText`  | Optional | N/A |
| `Run dynamic tests` | Enable to run runtime discoverable tests or dynamic tests  | Optional | `false`  |
| `Wait for build results` | Let pipeline wait for BrowserStack to complete the execution and get the test results  | Optional | `true` |
| `Test capabilities` | Enter capabilities in a key-value pair format on a new line. <br /><br />**For example**: <br />`coverage=true` <br />`geoLocation=CN"` | Optional | N/A |

</details>

<details>
<summary>Outputs</summary>

| Environment Variable | Description |
| --- | --- |
| `$BROWSERSTACK_BUILD_URL` |BrowserStack Dashboard url for the executed build|
| `$BROWSERSTACK_BUILD_STATUS`| Status of the executed build. Check out the [test results guide](https://www.browserstack.com/docs/app-automate/xcuitest/view-test-results) to learn about available status  |

</details>

## Troubleshooting

For internal troubleshooting, we would recommend that you start with the [troubleshooting guide](https://devcenter.bitrise.io/en/builds/build-data-and-troubleshooting.html).

If you are still unable to figure out the problem, please feel free to create an [issue](https://github.com/browserstack/browserstack-bitrise-xcui-step/issues), we will look into it ASAP.

## Contribution Guidelines

1. Fork this [repository](https://github.com/browserstack/browserstack-bitrise-xcui-step).
2. Add your changes.
3. Test your changes.
4. Raise a PR against this [repository](https://github.com/browserstack/browserstack-bitrise-xcui-step)
5. Work on comments, if any.
6. Once approved by our maintainers, we will merge the PR.
7. We will mention your name when we publish our [release](https://github.com/browserstack/browserstack-bitrise-xcui-step/releases) with your contribution. :slightly_smiling_face:
