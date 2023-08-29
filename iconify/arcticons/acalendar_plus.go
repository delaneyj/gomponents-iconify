package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcalendarPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.875 9.611h-2.056V6.528A1.028 1.028 0 0 0 34.792 5.5h-3.084a1.028 1.028 0 0 0-1.027 1.028V9.61H17.319V6.528A1.028 1.028 0 0 0 16.292 5.5h-3.084a1.028 1.028 0 0 0-1.027 1.028V9.61h-2.056a2.056 2.056 0 0 0-2.056 2.056v28.777a2.056 2.056 0 0 0 2.056 2.056h27.75a2.056 2.056 0 0 0 2.056-2.056V11.667a2.056 2.056 0 0 0-2.056-2.056Zm-1.542 6.681h-6.166"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.953 30.541a7.278 7.278 0 0 1-7.278 7.279h0a7.278 7.278 0 0 1-7.278-7.278V25.81a7.278 7.278 0 0 1 7.278-7.278h0a7.278 7.278 0 0 1 7.278 7.278m2.912 12.01a2.911 2.911 0 0 1-2.911-2.911V18.532m6.296-5.324v6.167"/>`),
		g.Group(children),
	)
}