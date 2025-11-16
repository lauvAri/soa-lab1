from flask import Flask, jsonify, request

app = Flask(__name__)

# 模拟数据库
borrows = [
    {"id": 1, "user_id": 101, "material_id": "1", "date": "2024-11-14"},
    {"id": 2, "user_id": 102, "material_id": "2", "date": "2024-11-15"},
]


@app.route("/borrows", methods=["GET"])
def get_borrows():
    return jsonify(borrows)


@app.route("/borrows", methods=["POST"])
def add_borrow():
    new_borrow = request.get_json()
    borrows.append(new_borrow)
    return jsonify(new_borrow), 201


if __name__ == "__main__":
    # 运行在 8081 端口
    app.run(port=8081, debug=True)
