設定ファイルから各ECSサービスのタスクのdesired countを０にします

```bash
takehito@deptop ~/stop-ecs-tasks (main)
❯ AWS_PROFILE=let-cd-playground ./stop-ecs-tasks -h config.json
Usage of ./stop-ecs-tasks:
  -config_file string
        stop configuration file path (default "config.json")
  -help
        output usage
```