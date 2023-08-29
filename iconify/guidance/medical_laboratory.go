package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalLaboratory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2 23.5h20M8.646 15a4.5 4.5 0 1 0 4.262-7.408M5.124 15a7.519 7.519 0 0 0 4.84 4.22M6.5 9.5L8 11m-6 2.5h9m4.258-8.258l.712-.712a1.81 1.81 0 0 0 0-2.56l-.94-.94a1.81 1.81 0 0 0-2.56 0L7 6.5l3.5 3.5l2.408-2.408m2.35-2.35a7.503 7.503 0 0 1-1.222 13.978m1.222-13.978l-2.35 2.35M9.964 19.22a3.235 3.235 0 0 1 4.072 0m-4.072 0a3.233 3.233 0 0 0-.858 1.069L8.5 21.5s-1 2-3 2h13c-2 0-3-2-3-2l-.606-1.211a3.234 3.234 0 0 0-.858-1.069M15.5 1.5L16.6.4"/>`),
		g.Group(children),
	)
}