package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M513 832q8-7 160-157t178-177q19-18 45.5-18t45.5 18.5t19 45.5t-19 45q-22 22-173 174T610 922q-68 67-159.5 91t-183-1T107 919T13 760t-1-181.5T104 420L452 76Q528 0 636 0t184.5 75.5T897 258t-76 182L535 723q-47 45-115 45t-115.5-45.5t-47.5-108T304 506q17-17 115-117.5T531 274q19-19 45.5-19t45 19t18.5 45.5t-18 45.5q-11 10-112.5 113T398 591q-9 9-9 22t9.5 22t22.5 9t23-9l285-286q92-91-.5-182.5T544 167L193 512q-69 69-68.5 161.5t68 159t159.5 67T513 832z"/>`),
		g.Group(children),
	)
}