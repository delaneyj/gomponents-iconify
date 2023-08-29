package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaUnlock0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaUnlock1" stroke-opacity="1" d="M 48.003124,207.99947 V 144.00668"/><path id="galaUnlock2" stroke-opacity="1" d="M 208.00205,207.99948 V 143.99989"/><path id="galaUnlock3" stroke-opacity="1" d="M 80.002888,239.99926 H 176.00225"/><path id="galaUnlock4" stroke-opacity="1" d="m 208.00205,207.99948 a 31.999787,31.999787 0 0 1 -31.9998,31.99978"/><path id="galaUnlock5" stroke-opacity="1" d="m 48.003124,207.99948 a 31.999787,31.999787 0 0 0 31.999764,31.99978"/><path id="galaUnlock6" stroke-opacity="1" d="m 128.00258,207.99949 v -31.9998"/><path id="galaUnlock7" d="M 47.996095,144.00668 A 15.999894,15.999894 0 0 1 63.995976,128.0068"/><path id="galaUnlock8" d="M 208.00205,143.99989 A 15.999894,15.999894 0 0 0 192.00218,128"/><path id="galaUnlock9" d="M 176.00225,64.00042 A 47.999683,47.999683 0 0 0 128.00258,16.00074"/><path id="galaUnlocka" d="M 80.002812,64.00042 A 47.999683,47.999683 0 0 1 128.00249,16.00074"/><path id="galaUnlockb" stroke-opacity="1" d="M 80.002888,128 V 63.907475"/><path id="galaUnlockc" stroke-opacity="1" d="m 176.00225,64.000424 0,15.999576"/><path id="galaUnlockd" stroke-opacity="1" d="M 64.003006,128 H 192.00218"/></g>`),
		g.Group(children),
	)
}