package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalWorkWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.822 30.098l.68-3.027l-2.33-2.05l1.582-2.67l-1.581-2.668l2.328-2.05l-.679-3.027l2.849-1.23l.289-3.09l3.089-.29l1.23-2.848l3.027.68l2.05-2.328l2.669 1.581L26.694 5.5l2.05 2.329l3.027-.679L33 9.998l3.089.29l.29 3.09l2.848 1.23l-.68 3.026l2.33 2.05l-1.582 2.67l1.581 2.669l-2.329 2.05l.68 3.027"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.024 8.684l-4.6 6.714l-7.807 2.301l4.963 6.451l-.224 8.136l7.668-2.726l7.669 2.727l-.224-8.137l4.963-6.451l-7.807-2.302ZM11.766 32.352s4.227 4.542 12.259 4.542s12.263-4.598 12.263-4.598M19.25 41.323a21.81 21.81 0 0 1-13.88-9.322l2.229.879l-.23-2.201s3.428 6.362 13.01 8.38c1.71.583.103 2.485-1.13 2.264Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.593 34.386l1.173-2.034s4.228 4.543 12.259 4.543s12.209-4.543 12.209-4.543"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.25 41.323a21.81 21.81 0 0 1-13.88-9.322l2.229.879l-.23-2.201s3.428 6.362 13.01 8.38c1.71.583.103 2.485-1.13 2.264Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.593 34.386l1.173-2.034s4.228 4.543 12.259 4.543"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.25 41.323a21.81 21.81 0 0 1-13.88-9.322l2.229.879l-.23-2.201s3.428 6.362 13.01 8.38c1.71.583.103 2.485-1.13 2.264Zm18.157-6.937l-1.173-2.034m-7.484 8.971A21.81 21.81 0 0 0 42.63 32l-2.228.88l.23-2.201s-3.428 6.362-13.01 8.38c-1.71.583-.104 2.485 1.13 2.264Z"/><circle cx="24.025" cy="21.654" r="4.788" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.379 39.058a18.279 18.279 0 0 0 7.242 0m-3.596-9.498V42.5m-4.244-8.413h8.487"/><circle cx="24.025" cy="21.654" r="7.726" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}