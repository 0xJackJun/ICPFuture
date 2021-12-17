HOW THIS PROJECT ORGANIZED ？
============================

The SRC folder is the blackhole canister, which MUST BE set to the controller of the canister you want to monitor. You can visit https://github.com/ninegua/ic-blackhole for more details. Since the cycle remaining in your canister is not public right now, you have to add blackhole to one of your controller so that the canister you moniting is public.

The SERVICE folder is the database and get_data canister, which your data is stored and exposed to others.
![Service-api](https://github.com/jackqr/ICPFuture/blob/main/image/service_api.png)

The HEARTBEAT folder is the cron job canister, which can do cron jobs on IC and calling the motoko function in SERVICE canister per hour.
![Heartbeat-api](https://github.com/jackqr/ICPFuture/blob/main/image/heartbeat_api.png)

The WATCH folder is off-chain program, which can do load balance, send emails and frond end. Since every canister has limit memorys, we can get how much data does canister store, and rebalance to each database canister.

The main logic is that heartbeat canister done two things: one is to call getData function in service canister, the other is to do cron job so it can consistantly get how much cycles left in the monitored canister. Since email service is not available on chian. So we have to do send email service off-chain. The off-chain services get the cycles data. And if it's less than the threshold user set originally, it will send email to the user.

TEST TESULT:
============
![Frontend](https://github.com/jackqr/ICPFuture/blob/main/image/frontend.jpeg)

![Email](https://github.com/jackqr/ICPFuture/blob/main/image/email_result.png)
