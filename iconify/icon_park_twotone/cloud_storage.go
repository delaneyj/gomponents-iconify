package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCloudStorage0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M44 29H4v13h40V29Z"/><path fill="#fff" d="M35.5 38a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M4 29L9.038 4.999H39.02l4.98 24"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19.006 16.026c-2.143 0-4.006 1.486-4.006 3.487C15 22 17.095 23 19.697 23h1.28m8.03-6.974c2.097 0 3.993.973 3.993 3.487C33 22 30.89 23 28.288 23h-1.3m2.019-6.974C29.007 13.042 27.023 11 24 11c-3.023 0-4.994 1.993-4.994 5.026"/><path stroke="#fff" stroke-width="4" d="M20 23h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCloudStorage0)"/>`),
		g.Group(children),
	)
}