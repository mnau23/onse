> Some notes for the future me, just in case.

## Variables

Below are explanations for some environment variables used in this project. See [.env.example](./../../.env.example) for reference.

**Case 1**: skip email sending during tests

If you do not want to send emails while testing, set:

```sh
DRY_RUN=true
```

**Case 2**: enable verbose logging

To show detailed logs during testing, set:

```sh
DEBUG_MODE=true
```

**Case 3**: test email rendering

To send test emails to your inbox, use:

```sh
DRY_RUN=false
DEBUG_MODE=true
```

If set, emails will be sent to the address specified in `EMAIL_RECEIVER_TEST`. This allows you to preview the actual email rendering and appearance.

**Case 4**: production run

For real execution (actual email sending, minimal logging), use:

```sh
DRY_RUN=false
DEBUG_MODE=false
```
