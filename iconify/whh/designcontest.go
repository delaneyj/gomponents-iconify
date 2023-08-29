package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Designcontest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M913 326L658 580L401 325q-15-15-36.5-15T328 325t-15 36.5t15 36.5l367 363q46 46 46 110t-46 109t-110.5 45T475 980L107 616q-69-69-94-162t0-185.5t94.5-162t163.5-94t187.5 0T622 107l36 36l47-46q21-21 57-29.5t78.5 2T913 110q47 47 47.5 109T913 326zm-240-5L545 193q-62-62-130-84t-129-4.5T173 174q-50 50-68 112.5t-3 122T161 513l384 384q32 32 64 0t0-64L238 461q-46-46-46-111t46-112q30-29 71.5-40t79-3.5T449 225l160 160q15 15 32 15t32-15q32-32 0-64z"/>`),
		g.Group(children),
	)
}