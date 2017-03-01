# sky_pinger


## How to deploy
apex deploy --env-file env.json

## Rule to call function daily

```
aws lambda add-permission --function-name skysports_sender --statement-id skysports_sender01 --action 'lambda:InvokeFunction' --principal events.amazonaws.com --source-arn arn:aws:lambda:eu-west-1:212733762802:rule/daily_at_noon
```

```
aws events put-rule --schedule-expression 'cron(0 12 * * ? *)' --name daily_at_noon
```

```
aws events put-targets --rule daily_at_noon --targets file://levels.json
```

