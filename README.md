## 实现功能：

 - 节点状态分为Leader（领导者）、Follower（追随者）、Candidate（候选人）
 - 节点间随机成为candidate状态并选举出Leader，且同时仅存在一个Leader
 - Leader节点定时发送心跳检测至其他Follower节点
 - Follower节点们超过一定时间未收到心跳检测，则Follower节点们重新开启选举
 - 客户端通过http发送消息到节点A，如果A不是Leader则转发至Leader节点
 - Leader收到客户端的消息后向Follower节点进行广播
 - Follower节点收到消息，反馈给Leader，等待Leader确认
 - Leader收到全网超过二分之一的反馈后，本地进行打印，然后将确认收到反馈的信息提交到Follower节点
 - Follower节点收到确认提交信息后，打印消息

### 说明：启动一个节点时，输入参数有三个，分别是程序名、节点id、节点端口号，如启动id为A的节点: raft.exe A 9000
