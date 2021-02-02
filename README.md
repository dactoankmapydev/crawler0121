# Crawler nguồn IOC
- IOCs (hashes, địa chỉ IP, tên miền…) được lấy từ các nhóm nội bộ đến các tổ chức, hoặc có thể từ đơn vị cung cấp thứ ba. Loại tìm kiếm này hầu như không chủ động nhưng lại mang về một số lợi ích trong quá trình tìm kiếm: nếu kẻ tấn công sử dụng lại cơ sở hạ tầng của phần mềm độc hại của chúng, sẽ rất dễ dàng để phát hiện ra các hoạt động độc hại ngay từ những bước đầu tiên, điều này đặc biệt đúng trong chiến dịch diện rộng, trong khi nó hiếm khi xảy ra với các mối đe doạ liên tục nâng cao (Advanced Persistent Threats) và các cuộc tấn công có chủ đích.

1. Crawler nguồn subscribed từ otx (người dùng subscribed tác giả nào thì nguồn bài viết sẽ hiện trên tài khoản người dùng
   đó)

- Trong 1 bài viết cụ thể có nhiều thông tin, có danh sách indicator ở bảng cuối cùng, cần lấy danh sách indicator để
  tích hợp vào ioc-profile, các thông tin còn lại lưu vào
  db [mẫu bài viết cụ thể](https://otx.alienvault.com/pulse/5fff646040d1907e50f04814)
- Đầu ra:

```
{
    "ioc_id": 2748114357,
    "ioc": "avsvmcloud.com",
    "ioc_type": "domain",
    "created_time": "2020-12-15T02:55:52",
    "crawled_time": "2021-01-13T13:59:23.206",
    "source": "otx-alienvault",
    "category": [
        "volexity",
        "dark halo",
        "solarwinds",
        "owa server",
        "exchange",
        "powershell",
        "adfind"
    ]
}
```

- Tiêu chí:

* [x]  Thu thập hết dữ liệu từ user AlientVault và các user được subcriber
* [x]  Có cơ chế phục hồi job khi job đột ngột stop
* [x]  Từ khi có nguy cơ mới đến khi thu thập < 1h
* [x]  Khi nguồn chết phải có thông báo
* [x]  Data crawler không bị duplicate

2. Thu thập nguồn compromised từ [mirror-h](https://mirror-h.org/) danh mục [archive](https://mirror-h.org/archive/)

- Đầu vào:
    - Thu thập các trường sau:
        - attacker (hostname)
        - country
        - web url (uid)
        - ip (src)
        - date (creation_date)
        - victim_hash=sha1(timestamp, uid, hostname)
        - timestamp=timestamp(creation_date)
- Đầu ra:

```
{
    "uid": "http://rednet.tech/",
    "hostname": "Noilesha",
    "src": "151.106.99.197",
    "victim_hash": "620cfe0d49455489e1a84b0b27f09b7206f7e90f",
    "creation_date": "2021-01-14T00:00:00",
    "timestamp": 1610557200,
    "country": "DE"
}
```

- Tiêu chí:

* [x]  Lấy được dữ liệu từ mirror-h bao gồm các thông tin: Domain bị chiếm điều khiển, IP chứa domain bị chiếm điều
  khiển, tên quốc gia chứa domain bị chiếm điều khiển, Kẻ tấn công, Thời gian cảnh báo
* [x]  Có cơ chế phục hồi job khi job đột ngột stop
* [x]  Từ khi có trang mới đến khi thu thập về không dưới 1h
* [x]  Khi nguồn chết phải có thông báo
* [x]  Data crawler không bị duplicate

3. Thu thập mẫu từ [virustotal](https://www.virustotal.com/api/v3/intelligence/hunting_notification_files)

- Đầu ra:

```
{
    'names': 'Order 2021-20073.doc',
    'sha256': 'dc9cc9c5768d1d2cc5fe315e941397e2cf42d666ce66abec36544e5e1da5a4d7',
    'sha1': '7c5afdf45731897b52b92e89fa3151544963de15',
    'md5': '54f6bf1f4626b877c39c766836b4d568',
    'tags': [
        'ole-embedded',
        'exploit',
        'rtf',
        'cve-2017-11882',
        'malware'
    ],
    'first_submit': '2021-01-18T02:27:50',
    'notification_date': '2021-01-18T03:28:36',
    'type_description': 'Rich Text Format',
    'magic': 'Rich Text Format data, version 1, unknown character set',
    'country': 'KR',
    'rule_name': 'Document_Malicious_Exploit',
    'detected': [
        'Antiy-AVL',
        'CAT-QuickHeal',
        'Cyren',
        'DrWeb',
        'Ikarus',
        'McAfee',
        'McAfee-GW-Edition',
        'NANO-Antivirus',
        'Qihoo-360',
        'Sangfor',
        'Symantec',
        'TrendMicro',
        'ZoneAlarm',
        'Zoner'
    ],
    'rate': 14,
    'point': 18
}
```

- Tiêu chí:

* [x]  Thu thập dữ liệu từ nguồn hunting của virustotal
* [x]  Tính điểm cho mẫu, lấy về nguồn trust
* [x]  Các dữ liệu thu thập được phải lưu trữ lại để phục vụ tra cứu sau này
* [x]  Có cơ chế phục hồi job khi job đột ngột stop
* [x]  Khi nguồn chết phải có thông báo
* [x]  Từ khi có nguy cơ mới đến khi thu thập < 1h
* [x]  Data crawler không bị duplicate
