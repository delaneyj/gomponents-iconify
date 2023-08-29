package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func S(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m280 220l45-47c-50-49-101-74-151-74c-43 0-78 14-107 41c-28 28-42 62-42 103c0 32 9 59 26 85c17 25 52 49 101 75c45 24 75 42 88 57c13 16 19 35 19 55c0 24-11 45-30 63c-19 17-43 27-71 27c-40 0-78-21-114-61L0 595c18 24 42 43 71 56s58 20 90 20c47 0 85-15 117-46c32-32 48-69 48-113c0-32-10-60-28-86c-18-25-55-51-107-78c-43-22-71-42-84-58s-20-33-20-50c0-20 9-38 25-53c16-14 35-22 58-22c36 0 73 18 110 55z"/>`),
		g.Group(children),
	)
}