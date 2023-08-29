package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tangerine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="39.073" r="21" fill="#f1b31c"/><circle cx="26" cy="46.346" r=".909"/><circle cx="26.909" cy="49.983" r=".909"/><circle cx="29.636" cy="47.255" r=".909"/><path fill="#e27022" d="M49.294 22.814a9.651 9.651 0 0 1-2.542 1.504a9.46 9.46 0 0 1-1.748.503a20.946 20.946 0 0 1-15.472 34.213a20.985 20.985 0 0 0 19.762-36.22Z"/><path fill="#b1cc33" d="M36 22.326a7.23 7.23 0 0 1 4.203.866c3.597 2.06 9.631.063 11.284-5.355a4.077 4.077 0 0 1-2.427-1.362c-2.48-2.987-10.247-3.212-13.06 5.851Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M36 22.326a7.23 7.23 0 0 1 4.203.866c3.597 2.06 9.631.063 11.284-5.355a4.077 4.077 0 0 1-2.427-1.362c-2.48-2.987-10.247-3.212-13.06 5.851Z"/><path d="M49.832 24.628a19.997 19.997 0 1 1-17.945-5.131"/><path d="M31.375 14.624s5.116 1.05 4.625 7.702"/></g>`),
		g.Group(children),
	)
}