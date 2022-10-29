arrivalRate=$1
# A local test that
#     1. Build the project
#     2. Spawn 5 replicas
#     3. Boot strap servers
#     4. Spawn a client
#     5. Compare the logs
#     6. Kill instances and clients

mage generate && mage build

raxos_path="replica/bin/replica"
ctl_path="client/bin/client"
output_path="logs/"

rm ${output_path}0.txt
rm ${output_path}0.log
rm ${output_path}1.txt
rm ${output_path}1.log
rm ${output_path}2.txt
rm ${output_path}2.log
rm ${output_path}3.txt
rm ${output_path}3.log
rm ${output_path}4.txt
rm ${output_path}4.log
rm ${output_path}5.txt
rm ${output_path}5.log

rm ${output_path}21.txt
rm ${output_path}21.log

rm ${output_path}local-test.log
rm ${output_path}status1.log
rm ${output_path}status2.log

echo "Removed old log files"

pkill replica
pkill replica
pkill replica
pkill replica
pkill replica
pkill client

echo "Killed previously running instances"

nohup ./${raxos_path} --name 0 --debugOn --debugLevel 0 >${output_path}0.log &
nohup ./${raxos_path} --name 1 --debugOn --debugLevel 0 >${output_path}1.log &
nohup ./${raxos_path} --name 2 --debugOn --debugLevel 0 >${output_path}2.log &
nohup ./${raxos_path} --name 3 --debugOn --debugLevel 0 >${output_path}3.log &
nohup ./${raxos_path} --name 4 --debugOn --debugLevel 0 >${output_path}4.log &

echo "Started 5 servers"

sleep 3

./${ctl_path} --name 21 --requestType status --operationType 1 >${output_path}status1.log

echo "Sent initial status to bootstrap"

sleep 3

echo "Starting client[s]"

nohup ./${ctl_path} --name 5 --debugOn --debugLevel 0  --requestType request  --arrivalRate "${arrivalRate}">${output_path}21.log &

sleep 200

echo "Completed Client[s]"

./${ctl_path} --name 21 --requestType status --operationType 2 >${output_path}status2.log

echo "Sent status to print log"

sleep 20

# python3 experiments/python/overlay-test.py ${output_path}0.txt ${output_path}1.txt ${output_path}2.txt ${output_path}3.txt ${output_path}4.txt >${output_path}local-test.log

pkill replica
pkill replica
pkill replica
pkill replica
pkill replica
pkill client

echo "Killed previously running instances"

echo "Finish test"