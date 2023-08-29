package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConeTestOnWallsNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsConeTestOnWallsNegative0)"><path d="M30.588 28.68C29.076 27.966 28 26.14 28 24c0-2.14 1.076-3.966 2.588-4.68l-15.013-4.504C16.446 17.018 17 20.316 17 24c0 3.685-.554 6.982-1.425 9.184l15.013-4.504Zm-16.657 3.173c-.318.953-.653 1.588-.931 1.948c-.278-.36-.613-.995-.931-1.948C11.429 29.935 11 27.159 11 24s.43-5.935 1.069-7.853c.318-.953.653-1.588.931-1.948c.278.36.613.995.931 1.948C14.571 18.065 15 20.841 15 24s-.43 5.935-1.069 7.853Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm17.354 37.578l-8.323 3.153A1.5 1.5 0 0 1 7 39.328V8.672A1.5 1.5 0 0 1 9.031 7.27l8.323 3.152a1 1 0 0 1 .646.935v2.099l14.08 4.224c.589-1.487 1.994-2.534 3.634-2.534c2.17 0 3.929 1.832 3.929 4.091c0 .214-.016.424-.046.63c1.434.841 2.403 2.442 2.403 4.28c0 1.836-.969 3.437-2.403 4.28c.03.204.046.414.046.628c0 2.26-1.759 4.091-3.929 4.091c-1.735 0-3.208-1.172-3.728-2.798L18 34.544v2.099a1 1 0 0 1-.646.935Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsConeTestOnWallsNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}