package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExerciseBicycleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsExerciseBicycleNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm37.215 36.954a9 9 0 0 0-3.508-16.926l-2.93-7.954l3.588-1.405a1 1 0 0 0 .574-.588l1-2.738a1 1 0 1 0-1.878-.686l-.846 2.314l-4.08 1.598a1 1 0 0 0-.573 1.277l3.043 8.261A9.001 9.001 0 0 0 25.515 24h7.062a5.423 5.423 0 0 1 1.43 10.655l-6.202 1.695A8.959 8.959 0 0 0 33 38a8.98 8.98 0 0 0 2.281-.292L37 40h-1a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2h-1.5l-2.285-3.046ZM36 29.42v.003a3.423 3.423 0 0 1-2.52 3.302l-.432.118l-17.375 4.965a6.013 6.013 0 0 1-2.188.212L12 40h1a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2h1.5l1.914-2.552C9.427 36.489 8 34.462 8 32.02c0-2.7 1.777-4.984 4.225-5.748L10.265 20H8.85a2.85 2.85 0 1 1 .284-5.687l4.502.45a2.625 2.625 0 0 1-.26 5.237H12.36l1.875 6H32.58A3.42 3.42 0 0 1 36 29.42ZM15.732 33A2 2 0 0 1 12 32a2 2 0 0 1 3.732-1H19a1 1 0 1 1 0 2h-3.268Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsExerciseBicycleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}