package fontelico

import (
	g "github.com/maragudk/gomponents"
)

const IconifyVersion = ""

func IconFromName(name string) g.Node {
	switch name {
	case "chrome": return Chrome()
	case "crown": return Crown()
	case "crownMinus": return CrownMinus()
	case "crownPlus": return CrownPlus()
	case "emoAngry": return EmoAngry()
	case "emoBeer": return EmoBeer()
	case "emoCoffee": return EmoCoffee()
	case "emoCry": return EmoCry()
	case "emoDevil": return EmoDevil()
	case "emoDispleased": return EmoDispleased()
	case "emoGrin": return EmoGrin()
	case "emoHappy": return EmoHappy()
	case "emoLaugh": return EmoLaugh()
	case "emoSaint": return EmoSaint()
	case "emoShoot": return EmoShoot()
	case "emoSleep": return EmoSleep()
	case "emoSquint": return EmoSquint()
	case "emoSunglasses": return EmoSunglasses()
	case "emoSurprised": return EmoSurprised()
	case "emoThumbsup": return EmoThumbsup()
	case "emoTongue": return EmoTongue()
	case "emoUnhappy": return EmoUnhappy()
	case "emoWink": return EmoWink()
	case "emoWinkTwo": return EmoWinkTwo()
	case "firefox": return Firefox()
	case "ie": return Ie()
	case "marquee": return Marquee()
	case "opera": return Opera()
	case "spinFive": return SpinFive()
	case "spinFour": return SpinFour()
	case "spinOne": return SpinOne()
	case "spinSix": return SpinSix()
	case "spinThree": return SpinThree()
	case "spinTwo": return SpinTwo()
	default:
		panic("Unknown icon name: " + name)
	}
}
