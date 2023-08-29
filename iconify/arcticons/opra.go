package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.882 29.348v-3.065h.996c.575 0 1.034.46 1.034 1.035s-.46 1.034-1.034 1.034h-.996m1.184-.017l.847 1.013m-3.838-.308a1.031 1.031 0 0 1-.74.308h0a1.03 1.03 0 0 1-1.035-1.034v-1.035c0-.575.46-1.034 1.035-1.034h0a.97.97 0 0 1 .713.301m-5.291 1.768c0 .575.46 1.034.996 1.034a1.03 1.03 0 0 0 1.035-1.034v-1.035c0-.574-.46-1.034-1.035-1.034s-.996.46-.996 1.034v1.035ZM17.516 16.45H14.1v2.733m3.416 11.377H14.1v-2.733m2.773-6.949h12.602m-12.602 1.544h12.602m-12.602 1.543h12.602m-12.602-4.858h12.602m-.43-2.758h3.416l.025 2.733M29.07 30.56h3.416v-2.733"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.08 22.346c-.17 7.258-4.582 13.115-8.326 13.943c-10.969 2.427-19.038-6.102-18.143-14.297c.809-7.406 9.869-10.882 18.295-9.54c7.63 1.214 8.345 2.535 8.174 9.894Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.009 8.447c10.793-.528 19.137 4.522 19.074 11.308c-.1 10.795-11.632 22.66-17.106 22.832c-5.363.168-17.88-13.162-18.931-23.297C4.242 11.539 14 8.838 22.009 8.447Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.274 8.79s-5.73-4.53-10.595-4.28c-4.296.22-5.34 5.693-5.34 5.693m28.735 9.958c-.127-.041 2.018 2.989 1.924 4.98c-.28 5.874-11.414 13.166-11.414 13.166m-8.616 4.175s-3.028 1.046-5.808 1.017c-5.975-.061-8.269-14.168-8.269-14.168"/>`),
		g.Group(children),
	)
}