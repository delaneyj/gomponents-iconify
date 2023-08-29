package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParamountPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.068 32.847l2.817 4.201h-4.3l-2.274-.84"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.027 19.318l-1.004 5.572l-2.372 5.634l5.486.247l-1.878 3.361v2.916h6.326"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.651 30.524l-.544 1.483l-1.384.84l-2.521 2.768l1.384-5.091l5.437-5.634"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.885 37.048h6.87l-7.957-8.946l-1.483-.148l-6.228-11.022l-7.809 9.045l-4.646 4.943l-4.794 6.128m19.473-10.761L24 24m18.154 2.683l-.958.972l-.78-1.097m1.26 2.375l-.48-1.278l-1.159.685m1.159-.685l1.401.298m-.451-6.695l-.633 1.21l-1.067-.822m1.9 1.904l-.833-1.082l-.908.994m.908-.994l1.427-.125m-2.389-6.271l-.251 1.342l-1.26-.473m2.373 1.265l-1.113-.792l-.578 1.217m.578-1.217l1.328-.536M37.51 11.58l.152 1.356l-1.343-.084m2.639.516l-1.296-.432l-.196 1.332m.196-1.332l1.114-.901m-5.488-3.863l.543 1.253l-1.31.313m2.675-.279l-1.365-.034l.201 1.332m-.201-1.332l.801-1.187m-6.377-2.09l.885 1.04l-1.161.681"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.455 6.821l-1.315.367l.582 1.214m-.582-1.214l.419-1.37m-6.709-.134L24 6.42l-.911.991m2.061-1.727L24 6.42l.911.991M24 6.42V4.987M5.846 26.683l.958.972l.78-1.097m-1.26 2.375l.48-1.278l1.159.685m-1.159-.685l-1.401.298m.451-6.695l.633 1.21l1.067-.822m-1.9 1.904l.833-1.082l.908.994m-.908-.994l-1.427-.125m2.389-6.271l.251 1.342l1.26-.473m-2.373 1.265l1.113-.792l.578 1.217M7.7 17.414l-1.328-.536m4.118-5.298l-.152 1.356l1.343-.084m-2.639.516l1.296-.432l.196 1.332m-.196-1.332l-1.114-.901m5.488-3.863l-.543 1.253l1.31.313m-2.675-.279l1.365-.034l-.201 1.332m.201-1.332l-.801-1.187m6.377-2.09l-.885 1.04l1.161.681"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.545 6.821l1.315.367l-.582 1.214m.582-1.214l-.419-1.37"/>`),
		g.Group(children),
	)
}