package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerDeck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m14.932 18.131l.027-4.261l-3.557 2.147zm2.131-4.262l-.021 4.256l3.557-2.147zm-.022 4.262v-.005l-.005.005zM31.995 6.12c0-1.172-.953-2.12-2.131-2.12H2.136A2.132 2.132 0 0 0 0 6.12v19.76C0 27.052.953 28 2.136 28h27.728A2.132 2.132 0 0 0 32 25.88V6.12zm-8.579 12.719h-.015c0 .781-.625 1.407-1.401 1.421l.047-.02H9.954a1.426 1.426 0 0 1-1.401-1.417h-.005v-5.651h.016c0-.771.615-1.401 1.375-1.412h12.079c.76.021 1.4.641 1.4 1.423v5.661z"/>`),
		g.Group(children),
	)
}