# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/notification.proto](#api_notification-proto)
    - [MarkNotificationsAsReadIn](#-MarkNotificationsAsReadIn)
    - [NewUserInvite](#-NewUserInvite)
    - [Notification](#-Notification)
    - [NotificationCountOut](#-NotificationCountOut)
    - [NotificationIn](#-NotificationIn)
    - [NotificationOut](#-NotificationOut)
    - [SendVerificationCodeIn](#-SendVerificationCodeIn)
  
    - [NotificationService](#-NotificationService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_notification-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/notification.proto



<a name="-MarkNotificationsAsReadIn"></a>

### MarkNotificationsAsReadIn
Входное сообщение для отметки уведомления как прочитанного


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notification_ids | [int64](#int64) | repeated | ID уведомлений |






<a name="-NewUserInvite"></a>

### NewUserInvite
Сообщение для отправки приглашения на платформу


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email нового пользователя |
| uuid | [string](#string) |  | UUID нового пользователя |






<a name="-Notification"></a>

### Notification
Сущность Уведомление


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID уведомления |
| text | [string](#string) |  | Тест уведомления |
| isRead | [bool](#bool) |  | состояние уведомления (прочитано или нет) |






<a name="-NotificationCountOut"></a>

### NotificationCountOut
Выходное сообщение для метода GetNotificationCount


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | Количество непрочитанных сообщений |






<a name="-NotificationIn"></a>

### NotificationIn
Входное сообщение для получения уведомлений


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | Лимит для уведомлений |
| offset | [int64](#int64) |  | Офсет для уведомлений |






<a name="-NotificationOut"></a>

### NotificationOut
ответсное сообщение на запрос уведомлений пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notifications | [Notification](#Notification) | repeated | Список сущностей Уведомление |






<a name="-SendVerificationCodeIn"></a>

### SendVerificationCodeIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email адресата |
| code | [string](#string) |  | Отправляемый код |





 

 

 


<a name="-NotificationService"></a>

### NotificationService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetNotificationCount | [.google.protobuf.Empty](#google-protobuf-Empty) | [.NotificationCountOut](#NotificationCountOut) | Метод получения количества непрочитанных уведомлений |
| GetNotification | [.NotificationIn](#NotificationIn) | [.NotificationOut](#NotificationOut) | Метод получения уведомлений по limit и offset |
| MarkNotificationsAsRead | [.MarkNotificationsAsReadIn](#MarkNotificationsAsReadIn) | [.google.protobuf.Empty](#google-protobuf-Empty) | Метод отметки уведомления как прочитанного |
| SendVerificationCode | [.SendVerificationCodeIn](#SendVerificationCodeIn) | [.google.protobuf.Empty](#google-protobuf-Empty) | Метод для отправки кода регистрации |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

