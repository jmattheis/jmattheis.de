package dns

import (
	"github.com/jmattheis/website/content"
	"github.com/miekg/dns"
	"strings"
)

type Config struct {
	Port string
}

func Listen(conf Config) {
	tty := &content.SingleText{
		Protocol:      "dns",
		Port:          conf.Port,
		Split:         ".",
		ForceBanner:   content.DnsSafeBanner,
		CommandPrefix: "dig @jmattheis.de +tcp +short ",
	}
	go func() {
		mux := dns.NewServeMux()
		mux.HandleFunc(".", func(w dns.ResponseWriter, msg *dns.Msg) {
			reply := new(dns.Msg)
			reply.SetReply(msg)

			question := ""
			if len(reply.Question) > 0 {
				question = strings.TrimSuffix(reply.Question[0].Name, ".")
			}

			exec := tty.Get(question)
			write(reply, exec)

			_ = w.WriteMsg(reply)
		})
		if err := dns.ListenAndServe(":"+conf.Port, "tcp", mux); err != nil {
			return
		}
	}()
}

func write(m *dns.Msg, s string) {
	parts := strings.Split(s, "\n")
	max := 0
	for _, part := range parts {
		if max < len(part) {
			max = len(part)
		}
	}
	for _, line := range parts {
		m.Answer = append(m.Answer, &dns.TXT{
			Txt: []string{"  " + padRight(line, max+2, " ")},
			Hdr: dns.RR_Header{
				Name:   ".",
				Rrtype: dns.TypeTXT,
				Class:  dns.ClassINET,
				Ttl:    0,
			}})
	}
}

func times(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}

func padRight(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}
