# PostgresSQL

## Overview
- Hỗ trợ nhiều kiểu dữ liệu, bao gồm cả việc tự định nghia dữ liệu.
- Nhiều kiểu index.
- Có MVCC.
- Có replication.

## Join
- Là cách để liên kết 2 bảng thoả mãn một điều kiện nào đó (thường là foreign key và primary key).
- Inner join trả về kết quả là các records có điều kiện thoả mãn từ 2 bảng.
- Left join trả về kết quả là các records thoả mãn điều kiện và tất cả các records từ bảng bên trái.
- Right join trả về kết quả là các records thoả mãn điều kiện và tất cả các records từ bảng bên phải.
- Full join trả về tất cả các records có trong 2 bảng.
- Syntax:
  ```
  SELECT * FROM table <LEFT/RIGHT/FULL> JOIN <điều kiện>
  ```

## Sub query:
- Exists: trả về true nếu sub query có kết quả hoặc ngược lại.
- expression/row-constructor IN/NOT IN : trả về true nếu expression nằm trong kết quả sub query hoặc ngược lại.
- expression/row-constructor operator ANY: so sánh một tập kết quả của subquery bằng cách sử dụng phép so sánh.
- expression/row_constructor operator ALL: so sánh mọi tập kết quả của subquery bằng cách sử dụng phép so sánh.

## Index: 
- Là một phương pháp đánh index làm cải thiện hiệu năng khi truy vấn, thường được sử dụng bằng các cấu trúc dữ liệu như tree, hashing,...
- Btree: balance tree, được sử dụng để đánh index cho id (int), được sử dụng cho các cột có dữ liệu thưa, so sánh kém.
- Hash: các giá trị trong cột được hash, chỉ sử dụng được phép so sánh bằng.
- Bit map index: được sử dụng cho các cột có dữ liệu dày đặc.
- Gist: sử dụng nhiều cấu trúc cây để hỗ trợ nhiều phép so sánh hơn.

## Partition: 
- Chia một bảng thành các bảng nhỏ hơn để thuận tiện lưu trữ.
- Phân vùng tăng khả năng lưu trữ, tăng hiệu năng.
- Có 3 kiểu phân vùng chính
  * Range partitioning: chia bảng theo miền giá trị liên tục.
  * List partitioning: chia theo miền giá trị rời rạc.
  * Hash partitioning: chia theo modulo, đảm bảo dữ liệu được chia đều.
- Syntax:
  ```
  CREATE TABLE <tên_bảng> PARTITION BY RANGE/LIST <tên_cột> <điều kiện>
  ```
  hoặc
  ```
  CREATE TABLE <tên_bảng> PARTITION BY HASH <tên_cột> PARTITIONS n
  ```

## Transaction:
- Đảm bảo tính tin cậy của một đoạn thao tác trên CSDL.
- Đảm bảo qui tắc ACID có trong SQL.
- Sử dụng MVCC để người dùng thao tác với version mới nhất.
- Khi các transaction không hoàn tất (lỗi), có thể rollback về trạng thái trước khi thực hiện transaction hoặc một SAVEPOINT đã được định nghĩa.
- Khởi tạo transaction:
  ```
  BEGIN
  <transaction>
  ```
- Chạy transaction
  ```
  COMMIT
  ```
- Huỷ transaction
  ```
  ROLLBACK
  ```
- Khai báo SAVEPOINT:
  ```
  BEGIN
  <transaction>
  SAVEPOINT sv1;
  <transaction>
  ROLLBACK TO sv1;
  COMMIT;
  ```
