package telegram

const (
	msgHelp        = "Отправьте мне название города и я сообщу, какая там сейчас температура и как она ощущается."
	msgStart       = "Привет! Я бот прогноза погоды.\n\n" + msgHelp
	msgUnknownCity = "Извините, я не знаю такого города."

	slowWind   = "Ветра почти нет ✅"
	middleWind = "Умеренный ветер 🤔"
	strongWind = "Ветер очень сильный, будьте осторожны, выходя из дома ❗️"
	storm      = "Шторм, на улицу лучше не выходить ❌"
)
