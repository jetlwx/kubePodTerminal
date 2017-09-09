#!/bin/bash
source /etc/profile

function javapid() {
pid=`ps -ef |grep java |grep -v grep |awk '{print $2}'`
echo $pid
}

function javaThreadnum() {
ps -eLf |grep java| wc -l
}

function printJstack() {
pid=`javapid`
 [ "$pid" != "" ] && jstack $pid
}

function showLogpath() {
keyword=$1
[ "${keyword}" != ""  ] && {
      logfile=`ls -lvt  /home/logs/tomcatlogs/$SVC | grep ${keyword}| head -n 1 |awk '{print $NF}'`
      logpath=/home/logs/tomcatlogs/$SVC/${logfile}
   echo $logpath  
}
}

#get the kubectl exec -it $POD -- tail -f xxxxx PID'S
function showtailPID() {
  ps -ef |grep 'tail -f -n 1000'|grep -v grep |awk '{print $2}'
}

#get the kubectl exec -it $pod -- /bin/bash PID'S
function showbashPID() {
  ps -ef |grep '/bin/bash'|grep -v grep | grep -v start_app.sh|grep  -v '/bin/sh' |grep -v 'java'|awk '{print $2}'
}
function showLog() {
keyword=$1
[ "${keyword}" != ""  ] && {
      logfile=`ls -lvt  /home/logs/tomcatlogs/$SVC | grep ${keyword}| head -n 1 |awk '{print $NF}'`
      logpath=/home/logs/tomcatlogs/$SVC/${logfile}
       cat ${logpath}
  }
}


function tailfLog() {
keyword=$1
[ "${keyword}" != ""  ] && {
      logfile=`ls -lvt  /home/logs/tomcatlogs/$SVC | grep ${keyword}| head -n 1 |awk '{print $NF}'`
      logpath=/home/logs/tomcatlogs/$SVC/${logfile}
       tailf -n 1000 ${logpath}
  }
}

function showTomcatConf(){
filename=$1
[ "${filename}" != "" ] &&  {
   cat /usr/local/tomcat/conf/$filename
 }
}

function sshd() {
  action=$1
   case ${action} in
     "start")
      /usr/sbin/sshd
      ;;
     "stop")
      kill -9 `ps -ef |grep /usr/sbin/sshd |grep -v grep | awk '{print $2}'`
      ;;
  esac
}

$@

