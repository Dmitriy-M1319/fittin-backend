package llm

const requestFormatString = `Есть следующие показатели результатов теста MMPI
ScaleResult(scale: '(Hs)', value: %d),
ScaleResult(scale: '(D)', value: %d),
ScaleResult(scale: '(Hy)', value: %d),
ScaleResult(scale: '(Pd)', value: %d),
ScaleResult(scale: '(Mf)', value: %d),
ScaleResult(scale: '(Pa)', value: %d),
ScaleResult(scale: '(Pt)', value: %d),
ScaleResult(scale: '(Sc)', value: %d),
ScaleResult(scale: '(Ma)', value: %d),
ScaleResult(scale: '(Si)', value: %d),
Интерпретируй эти результаты согласно следующему примеру (без форматирования в Markdown, кроме списков)
"Профиль показывает выраженные пики по шкалам 4 (Психопатия), 8 (Шизофрения) и 7 (Тревожность), что может указывать на:

    Импульсивность и склонность к конфликтам (шкала 4)
    Необычность мышления и восприятия (шкала 8)
    Высокий уровень тревоги и беспокойства (шкала 7)

Низкий показатель по шкале 0 (Интроверсия) говорит о высокой общительности.
Рекомендуется консультация специалиста для детального анализа."`
