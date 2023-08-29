package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Metrictime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 3v2m0 38v2m19.972-14.511l-1.902-.618m-6.902 9.5l1.175 1.618m7.629-23.478l-1.902.618M36.343 7.011l-1.175 1.618M12.832 39.371l-1.175 1.618M5.93 29.871l-1.902.618m0-12.978l1.902.618m5.727-11.118l1.175 1.618M24 25.5V35m14.266-6.365l-12.841-4.172m4.82-7.797a2.349 2.349 0 0 0 2.358 2.359a2.276 2.276 0 0 0 2.271-2.358v-2.359a2.331 2.331 0 0 0-2.27-2.358a2.406 2.406 0 0 0-2.359 2.358Zm4.105-3.843l-3.494 5.328"/><circle cx="20.812" cy="18.239" r=".873" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.126 18.413a2.74 2.74 0 0 0 2.096.612h.262a2.312 2.312 0 0 0 2.271-2.271h0a2.312 2.312 0 0 0-2.271-2.271h-2.358v-2.446h4.629m8.996 6.988a1.752 1.752 0 0 0 1.747-1.747h0a1.752 1.752 0 0 0-1.747-1.747h0a1.752 1.752 0 0 0 1.747-1.747h0a1.752 1.752 0 0 0-1.747-1.747m-2.882.612a2.972 2.972 0 0 1 2.184-.612h.698m-2.882 6.376a2.842 2.842 0 0 0 2.184.612h.698"/><circle cx="20.812" cy="13.61" r=".873" fill="currentColor"/>`),
		g.Group(children),
	)
}