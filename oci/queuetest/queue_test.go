package queuetest

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/limits"
	"github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
	"github.com/oracle/oci-go-sdk/v65/queue"
)

func init() {
	fmt.Println("Exec init function in packge queuetest.")
}

func CreatePrivateIp() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := core.CreatePrivateIpRequest{CreatePrivateIpDetails: core.CreatePrivateIpDetails{
		VnicId: common.String("ocid1.vnic.oc1.uk-london-1.abwgiljtvsbdqku5qj5uj2opcz5xys572xynflud5jkwgb2p5c24dpicws3q"),
		//		DefinedTags:   map[string]map[string]interface{}{"EXAMPLE_KEY_aPprR": map[string]interface{}{"EXAMPLE_KEY_JltY3": "EXAMPLE--Value"}},
		DisplayName: common.String("Second IP"),
		//		FreeformTags:  map[string]string{"EXAMPLE_KEY_PRR3Z": "EXAMPLE_VALUE_FDAAH9UuvkX0f4wUioP0"},
		//		HostnameLabel: common.String("EXAMPLE-hostnameLabel-Value"),
		IpAddress: common.String("10.0.0.188")}}

	// Send the request using the service client
	resp, err := client.CreatePrivateIp(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

/*
	func TestCreatePrivateIp(t *testing.T) {
		CreatePrivateIp()
	}
*/
func ListPrivateIpsFromVnic() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := core.ListPrivateIpsRequest{Limit: common.Int(59),
		//		Page:      common.String("EXAMPLE-page-Value"),
		//		SubnetId:  common.String("ocid1.test.oc1..<unique_ID>EXAMPLE-subnetId-Value"),
		//		VlanId:    common.String("ocid1.test.oc1..<unique_ID>EXAMPLE-vlanId-Value"),
		VnicId: common.String("ocid1.vnic.oc1.uk-london-1.abwgiljtvsbdqku5qj5uj2opcz5xys572xynflud5jkwgb2p5c24dpicws3q")}

	// Send the request using the service client
	pollingCnt := 5
	for idx := 0; idx < pollingCnt; idx++ {
		resp, err := client.ListPrivateIps(context.Background(), req)
		if err == nil {

			fmt.Println(resp)
			cnt := strings.Count(resp.String(), "10.0.0.188")
			fmt.Printf("Times: %d second\n", idx+1)
			if cnt > 0 {
				break
			}

		}
		fmt.Println(err)
		time.Sleep(time.Second * 1)

	}
}

/*
	func TestListPrivateIpsFromVnic(t *testing.T) {
		ListPrivateIpsFromVnic()
	}
*/
func AttachVnic() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := core.AttachVnicRequest{AttachVnicDetails: core.AttachVnicDetails{CreateVnicDetails: &core.CreateVnicDetails{SubnetId: common.String("ocid1.subnet.oc1.uk-london-1.aaaaaaaajm2ehw6hh2dmp2cazhx3ommoeoq6sgtpyl7ncesswkdoydddkeia"),
		//		DefinedTags:            map[string]map[string]interface{}{"EXAMPLE_KEY_x0loJ": map[string]interface{}{"EXAMPLE_KEY_yGdgS": "EXAMPLE--Value"}},
		//		FreeformTags:           map[string]string{"EXAMPLE_KEY_unosm": "EXAMPLE_VALUE_8hG0DNX3qu9YzKq2t4og"},
		//		HostnameLabel:          common.String("EXAMPLE-hostnameLabel-Value"),
		//		NsgIds:                 []string{"EXAMPLE--Value"},
		//		SkipSourceDestCheck:    common.Bool(false),
		//		VlanId:                 common.String("ocid1.test.oc1..<unique_ID>EXAMPLE-vlanId-Value"),
		AssignPrivateDnsRecord: common.Bool(true),
		AssignPublicIp:         common.Bool(false),
		DisplayName:            common.String("MyVnicCreate202308")},
		DisplayName: common.String("MyVnicAttach202308"),
		InstanceId:  common.String("ocid1.instance.oc1.uk-london-1.anwgiljti6bw44aca7hd53qy7mr2ogeiqeftrev322wijybjhgrn6v7sazfa")}}

	// Send the request using the service client
	resp, err := client.AttachVnic(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func TestAttachVnic(t *testing.T) {
	AttachVnic()
}

func ListVnicAttachments() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := core.ListVnicAttachmentsRequest{
		//		AvailabilityDomain: common.String("Uocm:LHR-AD-1"),
		CompartmentId: common.String("ocid1.tenancy.oc1..aaaaaaaambewmgb7iarfugubqrog6q3gmvd5ubqd6jyl5wfcy2esg3jfzs2a"),
		InstanceId:    common.String("ocid1.instance.oc1.uk-london-1.anwgiljti6bw44aca7hd53qy7mr2ogeiqeftrev322wijybjhgrn6v7sazfa"),
		Limit:         common.Int(29)}

	// Send the request using the service client
	pollingCnt := 10
	vnicAttachName := "MyVnicAttach202308"
	keyWordStart := 95
	for idx := 0; idx < pollingCnt; idx++ {
		resp, err := client.ListVnicAttachments(context.Background(), req)
		if err == nil {

			//			fmt.Println(resp)
			res := resp.String()
			pos := strings.Index(res, vnicAttachName)
			if pos >= keyWordStart {
				cnt := strings.Count(res[pos-keyWordStart:pos], "ATTACHED")
				if cnt > 0 {
					fmt.Println(res[pos-keyWordStart : pos+len(vnicAttachName)])
					break
				}
			}
			fmt.Printf("Elapsed time: %d seconds.\n", idx+1)

		}
		time.Sleep(time.Second * 1)
	}

}

func TestListVnicAttachments(t *testing.T) {
	ListVnicAttachments()
}

func ListLimitValues() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := limits.NewLimitsClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := limits.ListLimitValuesRequest{
		//AvailabilityDomain: common.String("EXAMPLE-availabilityDomain-Value"),
		//		OpcRequestId:  common.String("L1HSESYIHH7MESJWKBQ7<unique_ID>"),
		ServiceName: common.String("compute"),
		//		SortBy:        limits.ListLimitValuesSortByName,
		//		SortOrder:     limits.ListLimitValuesSortOrderDesc,
		CompartmentId: common.String("ocid1.tenancy.oc1..aaaaaaaambewmgb7iarfugubqrog6q3gmvd5ubqd6jyl5wfcy2esg3jfzs2a"),
		Limit:         common.Int(100),
		Name:          common.String("standard-e4-core-count"),
		//		Page:          common.String("EXAMPLE-page-Value"),
		ScopeType: limits.ListLimitValuesScopeTypeAd}

	// Send the request using the service client
	resp, err := client.ListLimitValues(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

/*
	func TestListLimitValues(t *testing.T) {
		ListLimitValues()
	}
*/
func ListLimitDefinitions() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := limits.NewLimitsClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := limits.ListLimitDefinitionsRequest{Limit: common.Int(100),
		Name:         common.String("standard-e4-core-count"),
		OpcRequestId: common.String("AIOWWQ30ZXZIWVLXPN"),
		//		Page:          common.String("EXAMPLE-page-Value"),
		//ServiceName: common.String("blockchain"),
		ServiceName: common.String("compute"),
		//		SortBy:        limits.ListLimitDefinitionsSortByDescription,
		//		SortOrder:     limits.ListLimitDefinitionsSortOrderDesc,
		CompartmentId: common.String("ocid1.tenancy.oc1..aaaaaaaambewmgb7iarfugubqrog6q3gmvd5ubqd6jyl5wfcy2esg3jfzs2a")}

	// Send the request using the service client
	resp, err := client.ListLimitDefinitions(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

/*
func TestListLimitDefinitions(t *testing.T) {
	ListLimitDefinitions()
}
*/

func UpdateBackendSet() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := networkloadbalancer.NewNetworkLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := networkloadbalancer.UpdateBackendSetRequest{BackendSetName: common.String("backendset_2023-0731-1153"),
		NetworkLoadBalancerId: common.String("ocid1.networkloadbalancer.oc1.uk-london-1.amaaaaaai6bw44aaqlftgh366umdegk2k7apkej7fclsm6zcve4ritu47xna"),
		UpdateBackendSetDetails: networkloadbalancer.UpdateBackendSetDetails{Backends: []networkloadbalancer.BackendDetails{networkloadbalancer.BackendDetails{Weight: common.Int(16),
			IpAddress: common.String("10.0.0.78"),
			IsBackup:  common.Bool(false),
			IsDrain:   common.Bool(false),
			IsOffline: common.Bool(false),
			Name:      common.String("backendServer1"),
			Port:      common.Int(9090)}},
			IsPreserveSource: common.Bool(false),
			Policy:           common.String("FIVE_TUPLE")}}

	// Send the request using the service client
	resp, err := client.UpdateBackendSet(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

/*
func TestUpdateBackendSet(t *testing.T) {

	UpdateBackendSet()

}
*/

func ListPrivateIps() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).
	/*
		req := core.ListPrivateIpsRequest{IpAddress: common.String("EXAMPLE-ipAddress-Value"),
			Limit:    common.Int(61),
			Page:     common.String("EXAMPLE-page-Value"),
			SubnetId: common.String("ocid1.test.oc1..<unique_ID>EXAMPLE-subnetId-Value"),
			VlanId:   common.String("ocid1.test.oc1..<unique_ID>EXAMPLE-vlanId-Value"),
			VnicId:   common.String("ocid1.test.oc1..<unique_ID>EXAMPLE-vnicId-Value")}
	*/
	req := core.ListPrivateIpsRequest{VnicId: common.String("ocid1.vnic.oc1.uk-london-1.abwgiljtetjvhgjinvfnkizdyq7kmqph3l2vw2ygporg32wp4nvljem5orza")}
	//req := core.ListPrivateIpsRequest{SubnetId: common.String("ocid1.subnet.oc1.ap-singapore-1.aaaaaaaa6amtw3rm5552qzxz6ijkwou3axeaeks7qsesv42xeoeujoja6gta")}
	// Send the request using the service client
	start := time.Now()

	resp, err := client.ListPrivateIps(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
	fmt.Println("time spent:", time.Since(start).Seconds())
}

/*
func TestListPrivateIps(t *testing.T) {
	ListPrivateIps()

}
*/

func GetVnic() {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := core.GetVnicRequest{VnicId: common.String("ocid1.vnic.oc1.ap-singapore-1.abzwsljrzlmmbdzxrfv722vuzjytvd6fzpfcgzasa2wdlp7vg4ly4l7viw4q")}

	// Send the request using the service client
	resp, err := client.GetVnic(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

/*
func TestGetVnic(t *testing.T) {
	GetVnic()

}
*/

func SendMessages() {
	client, err := queue.NewQueueClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	client.Host = "https://cell-1.queue.messaging.eu-frankfurt-1.oci.oraclecloud.com"
	id := rand.Intn(10000)
	requestID := fmt.Sprintf("%s%d", "OPC-Request-", id)
	messageContent1 := fmt.Sprintf("%s%d", "Message body : ", id)
	messageContent2 := fmt.Sprintf("%s%d", "Message body : ", id+1)
	messageContent3 := fmt.Sprintf("%s%d", "Message body : ", id+2)
	// Create a request and dependent object(s).
	req := queue.PutMessagesRequest{OpcRequestId: common.String(requestID),
		PutMessagesDetails: queue.PutMessagesDetails{Messages: []queue.PutMessagesDetailsEntry{queue.PutMessagesDetailsEntry{Content: common.String(messageContent1)}, queue.PutMessagesDetailsEntry{Content: common.String(messageContent2)}, queue.PutMessagesDetailsEntry{Content: common.String(messageContent3)}}},
		QueueId:            common.String("ocid1.queue.oc1.eu-frankfurt-1.amaaaaaag4yinciaiklmmiw722xgih7sxqkyj43z4obuv7k4opmim7du7v3a")}

	// Send the request using the service client
	resp, err := client.PutMessages(context.Background(), req)
	helpers.FatalIfError(err)
	// Retrieve value from the response.
	fmt.Println(resp)
}

func GetMessages() {
	client, err := queue.NewQueueClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	client.Host = "https://cell-1.queue.messaging.eu-frankfurt-1.oci.oraclecloud.com"
	queueOCID := "ocid1.queue.oc1.eu-frankfurt-1.amaaaaaag4yinciaiklmmiw722xgih7sxqkyj43z4obuv7k4opmim7du7v3a"

	// Create a request and dependent object(s).
	id := rand.Intn(10000)
	requestID := fmt.Sprintf("%s%d", "OPC-Request-", id)
	req := queue.GetMessagesRequest{Limit: common.Int(2),
		OpcRequestId:        common.String(requestID),
		QueueId:             common.String(queueOCID),
		TimeoutInSeconds:    common.Int(10),
		VisibilityInSeconds: common.Int(60)}

	// Send the request using the service client
	resp, err := client.GetMessages(context.Background(), req)
	helpers.FatalIfError(err)
	// Retrieve value from the response.
	//fmt.Println(resp)

	messages := resp.GetMessages.Messages
	for _, val := range messages {
		fmt.Printf("ID==%d\n", *val.Id)
		fmt.Printf("Receipt==%s\n", *val.Receipt)
		fmt.Printf("Content==%s\n", *val.Content)

		req := queue.DeleteMessageRequest{MessageReceipt: common.String(*val.Receipt),
			OpcRequestId: common.String(requestID),
			QueueId:      common.String(queueOCID)}

		// Send the request using the service client
		_, err := client.DeleteMessage(context.Background(), req)
		helpers.FatalIfError(err)

		// Retrieve value from the response.
		//fmt.Println(resp)
	}
}

/* func TestOCIQueue(t *testing.T){
	t.Log("=========Send Message======")
	SendMessages()

	t.Log("=========Get Message======")
	GetMessages()

} */

func deferfunc() {
	fmt.Println("similar to 'finally' statement in Java")
	if err := recover(); err != nil {
		fmt.Println("this is not a fatal error. Now recovered from ", err)
		fmt.Println("Start to release the used resouces")
	}
}

/*
	func TestDeferFunc(t *testing.T) {
		defer deferfunc()
		fmt.Println("Started...")
		panic("computing error")
		//os.Exit(-1)
	}
*/
type Programmer interface {
	WriteHelloWorld() string
}
type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() string {
	return "hello world! GO!"
}

func (gp *JavaProgrammer) WriteHelloWorld() string {
	return "hello world! Java!"
}

func WriteProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

/* func TestInterface(t *testing.T) {

	goP := new(GoProgrammer)
	WriteProgram(goP)
	javaP := new(JavaProgrammer)
	WriteProgram(javaP)
} */

type IntConv func() int

func SpentTime(inFunc IntConv) IntConv {
	return func() int {
		start := time.Now()
		ret := inFunc()
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func SlowFunc() int {
	time.Sleep(time.Second * 1)
	fmt.Println("func1 exec over")
	return 1
}
func SlowFunc2() int {
	time.Sleep(time.Second * 2)
	fmt.Println("func2 exec over")
	return 1
}

/* func TestFunc(t *testing.T){
	sp1 := SpentTime(SlowFunc)
	sp1()
	sp2 := SpentTime(SlowFunc2)
	sp2()
} */

func DoSomething(p interface{}) {
	/* 	if i, ok := p.(int); ok {
	   		fmt.Println("Integer: ", i)
	   		return
	   	}
	   	if s, ok := p.(string); ok {
	   		fmt.Println("String: ", s)
	   		return
	   	}
	   	fmt.Println("Unknow Type") */
	switch v := p.(type) {
	case int:
		fmt.Println("Integer: ", v)
	case string:
		fmt.Println("String: ", v)
	default:
		fmt.Println("Unknow Type")

	}

}

/*
	 func TestEmptyInterface(t *testing.T) {
		DoSomething(10)
		DoSomething("10")
	}
*/
var InvalidValueERR = errors.New("n is a number between 2 and 50")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 50 {
		return nil, InvalidValueERR
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

/*
	 func TestGetFibnacci(t *testing.T) {
		if v, err := GetFibonacci(51); err != nil {
			t.Error(err)
			if err == InvalidValueERR {
				fmt.Println("specified error")
			}
		} else {
			t.Log(v)
		}
	}
*/
func ProcessString() {
	s := "LifecycleState = DETACHED TimeCreated = 2023 - 08 - 17 14: 13: 54.793 + 0000 UTC DisplayName = vnicattachment20230817141353"
	idx := strings.Index(s, "vnicattachment20230817")
	ret := s[idx-95 : idx]
	fmt.Printf("idx:%d, ret=%s", idx, ret)
}

/*
func TestProcessString(t *testing.T) {
	ProcessString()
}
*/
