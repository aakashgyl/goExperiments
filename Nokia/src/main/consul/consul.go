package consul
//
//import (
//	"errors"
//	consulapi "github.com/hashicorp/consul/api"
//	log "github.com/sirupsen/logrus"
//	"strings"
//	"time"
//)
//
//const (
//	keynotfound = "Key not found"
//	keyislocked = "Key is already locked"
//)
//
////KV is used to form a key value pair
//type KV struct {
//	Key   string
//	Value string
//}
//
//var (
//	Consuladdress *consulapi.Config = &consulapi.Config{Address: "consul:8500"}
//	TTL           string            = Configs.TTL
//	LockDelay     time.Duration     = Configs.LockDelay
//)
//
////function CreateClient can be used to create a consul client
//func CreateClient() (*consulapi.Client, error) {
//	//client, err := consulapi.NewClient(consulapi.DefaultConfig())
//	client, err := consulapi.NewClient(Consuladdress)
//
//	return client, err
//}
//
////function  CreateKVClient is used to create a consul kv client
//func CreateKVClient() (*consulapi.KV, error) {
//	//client, err := consulapi.NewClient(consulapi.DefaultConfig())
//	client, err := consulapi.NewClient(Consuladdress)
//	kv := client.KV()
//	return kv, err
//
//}
//
////PutKVEntry is used to add a Key Value pair in the Consul KV Store
//func PutKVEntry(kventry KV) (bool, error) {
//	kv, _ := CreateKVClient()
//	entry := &consulapi.KVPair{Key: kventry.Key, Value: []byte(kventry.Value)}
//	_, err := kv.Put(entry, nil)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][PutKVEntry] "+entry.Key+" %v\n", err)
//		err = errors.New("Failed to Add Entry into Consul")
//		return false, err
//	}
//	return true, nil
//}
//
////GetKVEntry is used  to get the key value pair from consul
//func GetKVEntry(key string) (KV, error) {
//	kv, _ := CreateKVClient()
//	var entry KV
//	var errr error
//	pair, _, err := kv.Get(key, nil)
//	//fmt.Println("It come here",err)
//	if pair != nil {
//		log.Debug("[MOSS][BaicellEdge][consulIO] constructing KV struct inside GetKVEntry")
//		entry = KV{Key: pair.Key, Value: string(pair.Value)}
//	} else if err == nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][GetKVEntry] Key "+key+" not found %v\n", err)
//		errr = errors.New(keynotfound)
//
//	} else if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][GetKVEntry] Consul not Reachable %v\n", err)
//		errr = errors.New(INTERNAL_ERROR_CODE_124)
//
//	}
//
//	return entry, errr
//}
//
////CreateSession:creating session, ttl should be as time in seconds appended by s eg:"15s" and behavior eg:"delete" or "release"
//func CreateSession(behavior string) (string, error) {
//	client, _ := CreateClient()
//	var errr error
//	session := client.Session()
//	id, _, err := session.Create(&(consulapi.SessionEntry{TTL: TTL, Behavior: behavior, LockDelay: LockDelay}), nil)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateSession]  %v\n", err)
//		errr = errors.New("Consul Session creation failed")
//
//	}
//
//	return id, errr
//}
//
////CreateSessionAcquireLock:All integrated function to acquire lock,by just passing the key name.
//func CreateSessionAcquireLock(key string) (bool, string, error) {
//	log.Debug("[MOSS][BaicellEdge][consulIO][CreateSessionAcquireLock] Entered CreateSessionAcquireLock")
//	var sessionid string
//	var Err, sessionError, errr error
//	var status bool
//	kv, _ := CreateKVClient()
//
//	sessionid, sessionError = CreateSession("release")
//	if sessionError != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateSessionAcquireLock]  %v\n", sessionError)
//		errr = errors.New("Cound'nt create Session")
//	}
//	value, err := GetKVEntry(key)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateSessionAcquireLock]  %v\n", err)
//		errr = err
//	}
//	if err == nil {
//		status, _, Err = kv.Acquire(&consulapi.KVPair{Key: key, Session: sessionid, Value: []byte(value.Value)}, nil)
//		if status == false {
//			log.Errorf("[MOSS][BaicellEdge][consulIO][CreateSessionAcquireLock] "+"Key :"+key+" %v\n", Err)
//			errr = errors.New(keyislocked)
//		}
//		if Err != nil {
//			log.Errorf("[MOSS][BaicellEdge][consulIO][CreateSessionAcquireLock]  %v\n", Err)
//			errr = Err
//		}
//	}
//	return status, sessionid, errr
//}
//
////CreateRequestEntries:creating kv entries for request messages with ttl
//func CreateRequestEntries(msgid, requestFile string) (bool, string, error) {
//	log.Debug("[MOSS][BaicellEdge][consulIO][CreateRequestEntries] Entered CreateRequestEntries")
//	var Err, errr error
//	var status bool
//	keyname := "Request/" + msgid
//	kv, _ := CreateKVClient()
//
//	sessionid, sessionError := CreateSession("delete")
//	if sessionError != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateRequestEntries] Error is %v\n", sessionError)
//		errr = errors.New("Session Creation Failed")
//	}
//	status, _, Err = kv.Acquire(&consulapi.KVPair{Key: keyname, Session: sessionid, Value: []byte(requestFile)}, nil)
//	if status == false {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateRequestEntries] "+"Key :"+msgid+" %v\n", Err)
//
//		errr = errors.New(keyislocked)
//	}
//	if Err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateRequestEntries] "+"Key :"+msgid+" %v\n", Err)
//		errr = errors.New("Acquiring Lock Failed")
//	}
//
//	return status, sessionid, errr
//}
//
////GetRequestEnteries:to get  the message entry added into the consul
//func GetRequestEnteries(msgid string) (KV, error) {
//	keyname := "Request/" + msgid
//	pair, err := GetKVEntry(keyname)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][GetRequestEnteries] "+"Key :"+msgid+" %v\n", err)
//	}
//	return pair, err
//
//}
//
////AddRequestToQueue:Add an entry to the request message queue
//func AddRequestToQueue(neid, msgid string) (string, bool, error) {
//	log.Debug("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] Entered AddRequestToQueue")
//	var update bool
//	var err, errr, ERR error
//	var s []string
//	var updatedValue string
//
//	entry, getError := GetKVEntry(neid)
//	if getError != nil && getError.Error() == keynotfound {
//		update, ERR = PutKVEntry(KV{Key: neid, Value: msgid})
//		if ERR != nil {
//			log.Errorf("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] Error is %v\n", ERR)
//			errr = errors.New("Adding Entry Failed")
//		}
//
//	} else if getError != nil && getError.Error() == "Consul not Reachable" {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] Error is %v\n", err)
//		errr = errors.New("Consul not Reachable")
//	}
//	if getError == nil && entry.Value == "" {
//		csalstatus, csalsessid, csalerr := CreateSessionAcquireLock(neid)
//		if csalerr != nil && csalerr.Error() == keyislocked {
//			errr = csalerr
//		} else if csalstatus == true {
//			updatedValue = msgid
//			update, ERR = PutKVEntry(KV{Key: neid, Value: updatedValue})
//			if ERR != nil {
//				log.Errorf("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] Error is %v\n", err)
//				errr = errors.New("Adding Request Failed " + msgid)
//			}
//			if update == true {
//				relstatus, relerr := ReleaseLock(neid, csalsessid)
//				if relstatus == true {
//					log.Debug("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] releasing lock was successful")
//				} else if relstatus == false {
//					log.Debug("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] releasing lock Failed")
//
//				}
//				//log.Errorf("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] Error is %v\n", relerr)
//				if relerr != nil {
//					errr = relerr
//
//				}
//
//			}
//		}
//
//	} else if getError == nil && entry.Value != "" {
//		csalstatus, csalsessid, csalerr := CreateSessionAcquireLock(neid)
//		if csalerr != nil && csalerr.Error() == keyislocked {
//			errr = csalerr
//		} else if csalstatus == true {
//			s = strings.Split(entry.Value, ",")
//			s = append(s, msgid)
//			updatedValue = strings.Join(s, ",")
//			update, err = PutKVEntry(KV{Key: neid, Value: updatedValue})
//			if err != nil {
//				log.Errorf("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] Error is %v\n", err)
//				errr = errors.New("Adding Request Failed " + msgid)
//			}
//			if update == true {
//				relstatus, relerr := ReleaseLock(neid, csalsessid)
//				if relstatus == true {
//					log.Debug("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] releasing lock was successful")
//				} else if relstatus == false {
//					log.Debug("[MOSS][BaicellEdge][consulIO][AddRequestToQueue] releasing lock Failed")
//
//				}
//				if relerr != nil {
//					errr = relerr
//
//				}
//			}
//		}
//
//	}
//
//	return updatedValue, update, errr
//
//}
//
////PollRequestfromQueue:Pop a request from the queue
//func PollRequestfromQueue(neid string) (string, bool, error) {
//	log.Debug("[MOSS][BaicellEdge][consulIO][PollRequestfromQueue] Entered PollRequestfromQueue")
//	var popvalue string
//	var err, errr error
//	var update bool
//	entry, geterror := GetKVEntry(neid)
//	if geterror == nil {
//		if entry.Value == "" {
//			log.Error("[MOSS][BaicellEdge][consulIO][PollRequestfromQueue] Queue is Empty")
//			errr = errors.New("Queue is empty")
//		} else {
//			csalstatus, csalsessid, csalerr := CreateSessionAcquireLock(neid)
//			if csalerr != nil && csalerr.Error() == keyislocked {
//				log.Errorf("[MOSS][BaicellEdge][consulIO][PollRequestfromQueue]  %v\n", csalerr)
//				errr = csalerr
//			} else if csalstatus == true {
//				s := strings.Split(entry.Value, ",")
//				popvalue = s[0]
//				s = s[1:]
//				updatedValue := strings.Join(s, ",")
//				update, err = PutKVEntry(KV{Key: neid, Value: updatedValue})
//				if err != nil {
//					log.Errorf("[MOSS][BaicellEdge][consulIO][PollRequestfromQueue]  %v\n", err)
//					errr = errors.New("Updating the queue failed " + neid)
//				}
//				if update == true {
//					relstatus, relerr := ReleaseLock(neid, csalsessid)
//					if relstatus == true {
//						log.Debug("[MOSS][BaicellEdge][consulIO][PollRequestfromQueue] releasing lock was successful")
//					} else if relstatus == false {
//						log.Debug("[MOSS][BaicellEdge][consulIO][PollRequestfromQueue] releasing lock Failed")
//
//					}
//					if relerr != nil {
//						errr = relerr
//
//					}
//				}
//			}
//
//		}
//	} else if geterror != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][PollRequestfromQueue]  %v\n", geterror)
//		errr = err
//	}
//
//	return popvalue, update, errr
//}
//
////CreateConnectionURL:Adding  connection url for the NE
//func CreateConnectionURL(neid, URL string) (bool, error) {
//	keyname := "ConnectionURL/" + neid
//	update, err := PutKVEntry(KV{Key: keyname, Value: URL})
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateConnectionURL] "+"ConnectionURL for "+neid+" failed %v\n", err)
//		err = errors.New("Adding Connection URL Failed")
//
//	}
//	return update, err
//}
//
////GetConnectionURLfunc :to get the connection URL
//func GetConnectionURL(neid string) (KV, error) {
//	keyName := "ConnectionURL/" + neid
//	update, err := GetKVEntry(keyName)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][GetConnectionURL]  %v\n", err)
//		if err.Error() == keynotfound {
//			err = errors.New(NE_REGISTRATION_ERROR)
//		}
//	}
//	return update, err
//}
//
////CreateIPEntry:Adding   IPAddress  for NE
//func CreateIPEntry(neid, IPAddress string) (bool, error) {
//	keyname := "IPNEMAP/" + IPAddress
//	update, err := PutKVEntry(KV{Key: keyname, Value: neid})
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][CreateIPEntry] %v\n", err)
//		err = errors.New("Adding NE IP Failed")
//
//	}
//	return update, err
//}
//
////GetIPEntry:To fetch the ip address of the NE
//func GetIPEntry(IPAddress string) (KV, error) {
//	keyName := "IPNEMAP/" + IPAddress
//	update, err := GetKVEntry(keyName)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][GetIPEntry]  %v\n", err)
//		err = errors.New("NE IP Adress not found")
//
//	}
//	return update, err
//}
//
////ReleaseLock:function to release lock
//func ReleaseLock(key, id string) (bool, error) {
//	var errr error
//	getRes, getErr := GetKVEntry(key)
//	if getErr != nil {
//		errr = getErr
//	}
//	kv, _ := CreateKVClient()
//
//	relStatus, _, relErr := kv.Release(&consulapi.KVPair{Key: key, Value: []byte(getRes.Value), Session: id}, nil)
//	if relErr != nil {
//		errr = relErr
//	}
//	return relStatus, errr
//}
//
////AddToQueueandMessageEntries:function to add request to queue and  create message entries
//func AddToQueueandMessageEntries(s []string, neid string) (bool, []string, error) {
//	var status bool
//	var err error
//	var sessionid string
//	var sessid []string
//	updatedValue := strings.Join(s, ",")
//	_, _, er := AddRequestToQueue(neid, updatedValue)
//	if er != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][AddToQueueandMessageEntries]  %v\n", err)
//		err = errors.New("Failed to Update the Queue")
//
//	}
//	for i := 0; i < len(s); i++ {
//		status, sessionid, err = CreateRequestEntries(s[i], s[i])
//		sessid = append(sessid, sessionid)
//		log.Debug("[MOSS][BaicellEdge][consulIO] ", sessionid)
//		if err != nil {
//			log.Errorf("[MOSS][BaicellEdge][consulIO]AddToQueueandMessageEntries %v\n", err)
//			err = errors.New("Failed to create message entry")
//
//		}
//	}
//	return status, sessid, err
//}
//
////SessionInfo:function to give session details
//func SessionInfo(sessionid string) (*consulapi.SessionEntry, error) {
//	var errr error
//	client, _ := CreateClient()
//
//	session := client.Session()
//	sessioninfo, _, err := session.Info(sessionid, nil)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO]SessionInfo %v\n", err)
//		errr = err
//	}
//	return sessioninfo, errr
//}
//
////DeleteRequestfromQueue:function  to delete request queue entries
//func DeleteRequestfromQueue(neid string, msgidsl []string) (bool, error) {
//	log.Debug("[MOSS][BaicellEdge][consulIO][DeleteRequestfromQueue] Entered DeleteRequestfromQueue")
//	var err, errr error
//	var update bool
//	entry, geterror := GetKVEntry(neid)
//	if geterror == nil {
//		if entry.Value == "" {
//			log.Error("[MOSS][BaicellEdge][consulIO]DeleteRequestfromQueue Queue is Empty " + neid)
//			errr = errors.New("Queue is empty " + neid)
//		} else {
//			csalstatus, csalsessid, csalerr := CreateSessionAcquireLock(neid)
//			if csalerr != nil && csalerr.Error() == keyislocked {
//				log.Errorf("[MOSS][BaicellEdge][consulIO][DeleteRequestfromQueue]  %v\n", csalerr)
//				errr = csalerr
//			} else if csalstatus == true {
//				s := strings.Split(entry.Value, ",")
//				for j, _ := range msgidsl {
//					msgid := msgidsl[j]
//					for i, v := range s {
//						if v == msgid {
//							s = append(s[:i], s[i+1:]...)
//						}
//					}
//				}
//				updatedValue := strings.Join(s, ",")
//				update, err = PutKVEntry(KV{Key: neid, Value: updatedValue})
//				if err != nil {
//					log.Errorf("[MOSS][BaicellEdge][consulIO][DeleteRequestfromQueue] %v\n", err)
//					errr = errors.New("Updating the queue failed " + neid)
//				}
//				if update == true {
//					relstatus, relerr := ReleaseLock(neid, csalsessid)
//					if relstatus == true {
//						log.Debug("[MOSS][BaicellEdge][consulIO][DeleteRequestfromQueue] releasing lock was successful")
//					} else if relstatus == false {
//						log.Debug("[MOSS][BaicellEdge][consulIO][DeleteRequestfromQueue] releasing lock Failed")
//
//					}
//					if relerr != nil {
//						errr = relerr
//
//					}
//				}
//
//			}
//
//		}
//	} else if geterror != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][DeleteRequestfromQueue] %v\n", geterror)
//		errr = err
//	}
//
//	return update, errr
//}
//
//func DeleteEntry(key string) {
//	kv, _ := CreateKVClient()
//	_, err := kv.Delete(key, nil)
//	if err != nil {
//		log.Errorf("[MOSS][BaicellEdge][consulIO][DeleteEntry] %v\n", err)
//
//	}
//
//}
