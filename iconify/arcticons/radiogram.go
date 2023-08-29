package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.284 35.064a18.5 18.5 0 1 1 29.444-.013"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.69 36.226v-5.898a1.645 1.645 0 0 0-1.645-1.646h0l-26.103.037a1.65 1.65 0 0 0-1.645 1.645l-.017 5.898"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.764 31.903h8.198a1.27 1.27 0 0 1 1.27 1.27h0a1.27 1.27 0 0 1-1.27 1.272h-8.198a1.27 1.27 0 0 1-1.271-1.271h0a1.27 1.27 0 0 1 1.27-1.271Z"/><rect width="10.74" height="2.377" x="12.41" y="36.74" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.188"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.69 35.099V41a1.645 1.645 0 0 1-1.646 1.645h0l-26.102-.04a1.65 1.65 0 0 1-1.645-1.645l-.017-5.895m17.283-10.249a2.747 2.747 0 1 1 .01-.247m-15.572 4.12l9.83-3.867m-3.602 3.756l4.43-1.847"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.87 25.688a7.292 7.292 0 1 1 12.312 2.878"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.454 27.889a12.944 12.944 0 1 1 24.627.542"/><circle cx="31.363" cy="35.611" r="4.353" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}