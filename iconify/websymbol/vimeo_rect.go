package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimeoRect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 89v824q0 36-26 62t-62 26H88q-36 0-62-26T0 913V89q0-36 26-62T88 1h824q36 0 62 26t26 62zM880 254q0-93-112-93q-74 0-139 53t-89 126q12-2 23-2q8 0 15 1q18 2 32 7t22 19.5t8 38.5q0 43-38 119t-75 76q-19 0-36-19q-24-25-38-94.5T433 351t-31-117t-70-52q-37 0-80 26t-104.5 80.5T80 347v5q5 5 10.5 14.5t12 14T121 385q11 0 33-7t34-7q26 0 43 40q5 13 12.5 37.5T254 481q14 40 37 130l6.5 26l8 32l9.5 32.5l12.5 35l14 31.5l17.5 29.5l20.5 22l25 16.5l28.5 5q67 0 145-65t138.5-156t105-182.5T877 288q3-18 3-34z"/>`),
		g.Group(children),
	)
}