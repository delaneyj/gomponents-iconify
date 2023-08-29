package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MalariaPvMicroscopeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" fill-rule="evenodd" clip-path="url(#healthiconsMalariaPvMicroscopeNegative0)" clip-rule="evenodd"><path d="M9.693 21.994c1.18 1.527 4.627.84 7.698-1.533c3.072-2.374 4.606-5.536 3.426-7.063c-1.18-1.527-4.628-.84-7.7 1.533c-3.07 2.374-4.604 5.536-3.424 7.063Zm2.255-.494a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1-3a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1.068 1.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm17.371 7.535c3.287 1.985 5.151 4.918 4.165 6.552c-.987 1.634-4.451 1.35-7.738-.635c-3.286-1.984-5.15-4.917-4.164-6.551c.986-1.634 4.45-1.35 7.737.634ZM33.947 32a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3.546.906a.453.453 0 1 0 0-.906a.453.453 0 0 0 0 .906Zm1.817-9.366c1.75-.817 1.833-4.333.188-7.851c-1.646-3.519-4.398-5.707-6.148-4.89c-1.749.819-1.832 4.334-.187 7.853c1.646 3.518 4.398 5.707 6.147 4.889ZM29.948 19a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm1.7 2.4a.7.7 0 1 0 0-1.4a.7.7 0 0 0 0 1.4Zm-11.08 9.218c1.784 3.55 1.8 7.146.035 8.033c-1.764.887-4.64-1.272-6.425-4.82c-1.784-3.55-1.8-7.146-.036-8.033c1.765-.887 4.642 1.271 6.426 4.82ZM18.948 35a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm.453-2.094a.453.453 0 1 0 0-.906a.453.453 0 0 0 0 .906Zm.947 3.794a.7.7 0 1 1-1.4 0a.7.7 0 0 1 1.4 0Z"/><path d="M48 0H0v48h48V0ZM24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/></g><defs><clipPath id="healthiconsMalariaPvMicroscopeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}