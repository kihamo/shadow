��          �   %   �      p  O   q     �  v   �     H     W     v     �  C   �     �  �   �     �  -   �     �               +     9     Q     a     p  i   �  *   �          &     8  v   M     �  ?  �  i        �  �   �     }	  #   �	     �	     �	  c   �	  %   ;
  w  a
     �  O   �     6     L     g     {  %   �     �     �     �  �   �  /   �     �  #   �  %   �  �        �                                                          	                                                
                 configAddress of jaeger-agent's HTTP sampling server (only for remote sampler) configConstant configControls how often the remotely controlled sampler will poll jaeger-agent for the appropriate sampling strategy configEnabled configEnabled metrics for RPC configEndpoint configHost configHow often the buffer is force-flushed, even if it's not full configLocal collector configMaximum number of operations that the sampler will keep track of. If an operation is not tracked, a default probabilistic sampler will be used rather than the per operation specific sampler configOthers configParam is a value passed to the sampler configPort number configProbabilistic configRate limiting configRemote configRemote collector configReporter configSampler configService scope configSize of internal queue where reported spans are stored before they are processed in the background configTags in format tag1_name=tag1_value configType configUser login configUser password configWhen enabled, enables LoggingReporter that runs in parallel with the main reporter and logs all submitted spans config-tabtracing Report-Msgid-Bugs-To: dev@kihamo.ru
Last-Translator: Kihamo Muramodo <dev@kihamo.ru>
Language-Team: 
Language: Russian
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 
Plural-Forms: nplurals=3; plural=n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2; Адрес HTTP сепмлер сервера (только для удаленного сэмплера) Постоянный Контролирует, как часто дистанционно управляемый сэмплер будет опрашивать jaeger агента для соответствующей стратегии выборки Включено Включены RPC метрики Конечная точка Хост Как часто буфер сбрасывается, даже если он не заполнен Локальный коллектор Максимальное количество операций, которые будет отслеживать сэмплер. Если операция не отслеживается, будет использоваться вероятностный сэмплер по умолчанию, а не конкретный сэмплер для каждой операции Другое Параметр это значение переданное в семплер Номер порта Вероятностный Рейт лимит Удаленный Удаленный коллектор Репортер Семплер Область сервиса Размер внутренней очереди, где спаны хранятся до их обработки в фоновом режиме Теги в формате tag1_name=tag1_value Тип Логин пользователя Пароль пользователя Когда включено, включает LoggingReporter, который работает параллельно с основным репортером и регистрирует все отправленные спаны Трейсинг 