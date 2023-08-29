package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hyperion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.374 27.269c-.64-.865-.922-1.151-2.134.316s-2.121-.107-2.778-2.56l-1.34-5l1.944-1.955A10.653 10.653 0 0 1 23.98 7.059a10.654 10.654 0 0 1 5.915 11.011l1.944 1.955l-1.34 5c-.657 2.453-1.566 4.028-2.778 2.56s-1.494-1.18-2.134-.316s-.956 1.294-1.607 1.294s-.966-.43-1.606-1.294Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.458 26.557s.035 5.604-2.25 7.742c0 0-4.035-.897-4.5 2.764h-2.111a1.314 1.314 0 0 0-1.388-.801h-2.75c-.275.013-.758.115-1.135 1.307m19.41-10.931a12.976 12.976 0 0 0 1.551 6.483s5.219-1.85 6.277 3.995a1.725 1.725 0 0 0 1.947 0s.1-.879.879-.879h2.96a.735.735 0 0 1 .785.752"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.935 27.993s.47 5.63-2.036 9.277h-1.564a1.962 1.962 0 0 0-2.066 2.05c-1.441.644-1.098-.68-2.413-.716c-1.26 0-2.987-.163-2.987 1.955c0 0-1.361-.427-1.405.144m14.702-12.878s.137 5.649 1.695 8.266h2.995a1.924 1.924 0 0 1 1.874 1.874c0 1.774 1.033 1.29 1.033 1.29s-.192-.632 1.736-.632s2.548.955 2.548 1.892l.702.014M19.95 25.405a49.532 49.532 0 0 1-1.563-5.143m9.624 5.143a49.532 49.532 0 0 0 1.563-5.143"/><circle cx="23.981" cy="14.405" r="1.962" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}