# vim: ts=4 sw=4 sts=4 noet
@smoke
Feature: Quick smoke tests

	Background:
		Disable terminal multiplexing for these tests

		Given I set the environment variable "CHELL_MUXING" to "false"

	Scenario: Chell is an interactive shell
		Given I run `chell` interactively

		When I type "echo foo"
		And I close the stdin stream

		Then stderr should not contain anything
		And stdout should contain "foo"
		And the exit status should be 0
