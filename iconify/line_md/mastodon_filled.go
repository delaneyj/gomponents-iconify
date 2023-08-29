package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MastodonFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdMastodonFilled0"><path fill="#fff" d="M17.25 2.315c2.366.346 4.361 2.131 4.67 4.396c.167 1.683.022 4.421.02 4.87c0 .133-.019 1.34-.027 1.468c-.207 3.236-2.247 4.514-4.391 4.922c-.03.008-.063.014-.097.02c-1.36.263-2.815.333-4.197.372c-.33.008-.66.008-.99.008c-1.373 0-2.742-.16-4.077-.479a.046.046 0 0 0-.059.047c.038.43.133.853.281 1.259c.185.47.832 1.597 3.234 1.597a17.89 17.89 0 0 0 4.146-.479a.048.048 0 0 1 .053.025a.047.047 0 0 1 .005.02v1.589a.05.05 0 0 1-.02.038c-.444.318-1.048.5-1.562.661c-.228.071-.459.133-.692.187c-2.12.478-4.331.362-6.388-.333c-1.921-.667-3.882-2.302-4.367-4.266a22.953 22.953 0 0 1-.545-3.233c-.151-1.64-.164-3.282-.229-4.928c-.045-1.148-.019-2.4.226-3.528c.51-2.292 2.61-3.896 4.91-4.233c.4-.058 1.15-.27 4.655-.27h.026c3.503 0 5.016.212 5.415.27Z"/><g fill="none" stroke="#000" stroke-dasharray="18" stroke-dashoffset="18" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.2 12L12.2 9C12.2 8 13.1 6.5 14.85 6.5C16.6 6.5 17.5 8 17.5 9L17.5 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="18;0"/></path><path d="M6.9 14L6.9 9C6.9 8 7.8 6.45 9.54 6.5C11.3 6.47 12.2 8 12.2 9L12.2 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMastodonFilled0)"/>`),
		g.Group(children),
	)
}