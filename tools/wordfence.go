/* https://github.com/blackcrw/akumascan/projects/1#card-70649811 */

package tools

import (
	"fmt"
	"strings"

	"github.com/blackcrw/akumascan/pkg/nettools"
)

func wordfence(channel chan [3]string, quit chan bool, response *nettools.Response) {
	var ( confidence int; blocked bool )

	var (
		snippet1 = "Message Exists in Source Code: \"Generated by Wordfence\""
		snippet2 = "Message Exists in Source Code: \"This response was generated by Wordfence\""
		blocked1 = "There is a chance the WAF has blocked you. Message: \"A potentially unsafe operation has been detected in your request to this site\""
		blocked2 = "WAF has blocked your access to the site. Message: \"Your access to this site has been limited\""
	)

	channel<-[3]string{"Wordfence", "", "Checking if "+snippet1}
	if strings.Contains(response.Raw, "Generated by Wordfence") {
		confidence += 4
		channel<-[3]string{"Wordfence", fmt.Sprint(confidence), snippet1}
	}
	
	channel<-[3]string{"Wordfence", "", "Checking if "+snippet2}
	if strings.Contains(response.Raw, "This response was generated by Wordfence") {
		confidence += 6
		channel<-[3]string{"Wordfence", fmt.Sprint(confidence), snippet2}
	}
	
	if strings.Contains(response.Raw, "A potentially unsafe operation has been detected in your request to this site") {
		confidence += 3
		channel<-[3]string{"Wordfence", "", blocked1}
	}
	
	if strings.Contains(response.Raw, "Your access to this site has been limited") {
		confidence += 4
		channel<-[3]string{"Wordfence", "", blocked2}
	}

	channel<-[3]string{"Wordfence", fmt.Sprint(confidence), fmt.Sprint(blocked)}


	quit <- true
}