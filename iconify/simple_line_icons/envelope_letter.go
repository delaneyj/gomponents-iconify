package simple_line_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeLetter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path d="M1023 473q0-4-1-8t-2.5-8t-3.5-7.5t-5-6.5l-1-1q-2-3-5-6l-6-6l-167-168V159q0-13-9.5-22.5T800 127H699L582 27q-11-11-24.5-17.5T529 1t-30 0t-28.5 8.5T446 27L330 127H224q-13 0-22.5 9.5T192 159v105L29 427Q0 446 0 481v479q0 27 19 45.5t45 18.5h896q17 0 32-8.5t23.5-23.5t8.5-32V478l-1-5zM269 726L64 916V545zm62 30q6-3 11-9q1-1 2-2.5t2-3.5l138-128q11-9 25-9t24 8l384 348H110zm427-26l202-179v361zm140-313h-1l55 55l-120 107V351zM491 72q10-9 23-9q6 0 12 2.5t11 6.5l64 55H427zm277 119v445l-57 51l-137-124q-6-5-13-9t-14-6.5t-14.5-4.5t-15.5-2.5t-15.5 0T486 543t-15 5t-14.5 7t-13.5 9L316 682l-60-53V191h512zM104 443l88-88v217L75 469l27-26h2z" fill="currentColor"/>`),
		g.Group(children),
	)
}